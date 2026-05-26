import {
    fetchContractorAnalytics,
    fetchContractorNames,
    fetchContractorObjectDetails
} from '../api/contractorAnalyticsApi.js';

function escapeHtml(value) {
    return String(value)
        .replaceAll('&', '&amp;')
        .replaceAll('<', '&lt;')
        .replaceAll('>', '&gt;')
        .replaceAll('"', '&quot;')
        .replaceAll("'", '&#39;');
}

function formatMoney(value) {
    const amount = Number(value) || 0;
    return `${new Intl.NumberFormat('ru-RU', { maximumFractionDigits: 0 }).format(amount)} ₽`;
}

function formatPercent(value, fractionDigits = 0) {
    if (value === null || value === undefined) return '—';
    const num = Number(value);
    if (Number.isNaN(num)) return '—';
    return `${num.toFixed(fractionDigits)}%`;
}

function formatDate(value) {
    if (!value) return '—';
    const date = new Date(value);
    if (Number.isNaN(date.getTime())) return '—';
    return date.toLocaleDateString('ru-RU');
}

function formatReadiness(value) {
    if (value === null || value === undefined) return 'Нет данных';
    const text = String(value).trim();
    return text === '' ? 'Нет данных' : text;
}

function parseReadinessToPercent(value) {
    if (value === null || value === undefined) return null;
    const normalized = String(value).trim().replaceAll(',', '.').replaceAll('%', '');
    if (!normalized) return null;

    const parsed = Number(normalized);
    if (Number.isNaN(parsed)) return null;

    return Math.max(0, Math.min(100, parsed));
}

function riskClass(level) {
    switch (level) {
    case 'ok':
        return 'ca-risk-ok';
    case 'risk':
        return 'ca-risk-risk';
    case 'critical':
        return 'ca-risk-critical';
    default:
        return 'ca-risk-no-data';
    }
}

function statusMeta(details) {
    const overdueDebt = Number(details?.overdue_debt_amount) || 0;
    const endDate = details?.work_end_date ? new Date(details.work_end_date) : null;
    const isOverdueByDates = overdueDebt > 0 && endDate && !Number.isNaN(endDate.getTime()) && endDate < new Date();
    if (isOverdueByDates) {
        return { text: 'Просрочен', className: 'ca-status-overdue' };
    }

    const raw = String(details?.status || '').toLowerCase();
    if (raw.includes('проср') || raw.includes('рїс')) return { text: 'Просрочен', className: 'ca-status-overdue' };
    if (raw.includes('работ') || raw.includes('сђр°р±')) return { text: 'В работе', className: 'ca-status-active' };
    return { text: 'Нет данных', className: 'ca-status-unknown' };
}

export class ContractorAnalytics {
    constructor(container) {
        this.container = container;
        this.contractors = [];
        this.selectedContractor = '';
        this.analytics = null;
        this.emptyStateMessage = '';
        this.loading = false;
        this.loadingContractors = false;
        this.error = null;
        this.requestToken = 0;
        this.connectionFrame = 0;
        this.connectionMarkerSeq = 0;
        this.handleResize = () => this.scheduleConnectionsRender();
        window.addEventListener('resize', this.handleResize);
        this.init();
    }

    async init() {
        this.loadingContractors = true;
        this.render();

        try {
            const names = await fetchContractorNames();
            this.contractors = Array.isArray(names) ? names : [];
        } catch (err) {
            this.error = `Не удалось загрузить список подрядчиков: ${err.message}`;
        } finally {
            this.loadingContractors = false;
            this.render();
        }
    }

    async loadAnalytics(contractorName) {
        if (!contractorName) return;

        const token = ++this.requestToken;
        this.loading = true;
        this.error = null;
        this.emptyStateMessage = '';
        this.selectedContractor = contractorName;
        this.render();

        try {
            const data = await fetchContractorAnalytics(contractorName);
            if (token !== this.requestToken) return;

            this.analytics = {
                ...data,
                selected_object: null
            };
            this.emptyStateMessage = '';
        } catch (err) {
            if (token !== this.requestToken) return;
            const message = String(err?.message || '');

            if (message.includes('contractor has no objects')) {
                this.emptyStateMessage = 'По выбранному подрядчику нет объектов (construction_object).';
                this.error = null;
                this.analytics = null;
                return;
            }

            this.error = `Не удалось загрузить аналитику: ${message}`;
            this.analytics = null;
        } finally {
            if (token === this.requestToken) {
                this.loading = false;
                this.render();
            }
        }
    }

    async loadObjectDetails(customerName, objectName) {
        if (!this.selectedContractor || !customerName || !objectName) return;

        const token = ++this.requestToken;
        this.loading = true;
        this.error = null;
        this.render();

        try {
            const details = await fetchContractorObjectDetails(this.selectedContractor, customerName, objectName);
            if (token !== this.requestToken) return;
            if (this.analytics) {
                this.analytics.selected_object = details;
            }
        } catch (err) {
            if (token !== this.requestToken) return;
            this.error = `Не удалось загрузить объект: ${err.message}`;
        } finally {
            if (token === this.requestToken) {
                this.loading = false;
                this.render();
            }
        }
    }

    scheduleConnectionsRender() {
        if (this.connectionFrame) {
            cancelAnimationFrame(this.connectionFrame);
        }

        this.connectionFrame = requestAnimationFrame(() => {
            this.connectionFrame = 0;
            this.renderConnections();
        });
    }

    renderConnections() {
        const groups = this.container.querySelectorAll('[data-ca-group]');
        if (groups.length === 0) return;

        const namespace = 'http://www.w3.org/2000/svg';
        groups.forEach((group) => {
            const linksLayer = group.querySelector('[data-ca-links]');
            const customerNode = group.querySelector('[data-ca-customer-node]');
            const objectNodes = group.querySelectorAll('[data-ca-object-node]');

            if (!linksLayer || !customerNode || objectNodes.length === 0) {
                if (linksLayer) linksLayer.textContent = '';
                return;
            }

            const groupRect = group.getBoundingClientRect();
            const customerRect = customerNode.getBoundingClientRect();
            const width = Math.max(1, Math.ceil(groupRect.width));
            const height = Math.max(1, Math.ceil(groupRect.height));

            linksLayer.setAttribute('viewBox', `0 0 ${width} ${height}`);
            linksLayer.setAttribute('width', String(width));
            linksLayer.setAttribute('height', String(height));
            linksLayer.textContent = '';

            const markerID = `ca-link-arrow-${this.connectionMarkerSeq++}`;
            const arrowSize = 5.4;
            const defs = document.createElementNS(namespace, 'defs');
            const marker = document.createElementNS(namespace, 'marker');
            marker.setAttribute('id', markerID);
            marker.setAttribute('markerWidth', String(arrowSize));
            marker.setAttribute('markerHeight', String(arrowSize));
            marker.setAttribute('refX', String(arrowSize - 0.2));
            marker.setAttribute('refY', String(arrowSize / 2));
            marker.setAttribute('orient', 'auto');
            marker.setAttribute('markerUnits', 'strokeWidth');

            const markerPath = document.createElementNS(namespace, 'path');
            markerPath.setAttribute('d', `M0,0 L${arrowSize},${arrowSize / 2} L0,${arrowSize} z`);
            markerPath.setAttribute('fill', '#60a5fa');
            marker.appendChild(markerPath);
            defs.appendChild(marker);
            linksLayer.appendChild(defs);

            const startXBase = customerRect.right - groupRect.left;
            const startYBase = customerRect.top - groupRect.top + customerRect.height / 2;

            objectNodes.forEach((objectNode) => {
                const objectRect = objectNode.getBoundingClientRect();
                const startX = startXBase;
                const startY = startYBase;
                const endX = objectRect.left - groupRect.left;
                const endY = objectRect.top - groupRect.top + objectRect.height / 2;
                const deltaX = endX - startX;
                const bend = Math.max(24, Math.abs(deltaX) * 0.45);

                const control1X = startX + bend;
                const control1Y = startY;
                const control2X = deltaX >= 0 ? endX - bend : endX + bend;
                const control2Y = endY;

                const path = document.createElementNS(namespace, 'path');
                path.setAttribute('d', `M ${startX} ${startY} C ${control1X} ${control1Y}, ${control2X} ${control2Y}, ${endX} ${endY}`);
                path.setAttribute('class', 'ca-link-path');
                path.setAttribute('marker-end', `url(#${markerID})`);
                linksLayer.appendChild(path);
            });
        });
    }

    bindEvents() {
        const select = this.container.querySelector('[data-ca-select]');
        if (select) {
            select.addEventListener('change', (event) => {
                const contractorName = event.target.value;
                if (!contractorName) {
                    this.selectedContractor = '';
                    this.analytics = null;
                    this.emptyStateMessage = '';
                    this.error = null;
                    this.loading = false;
                    this.requestToken += 1;
                    this.render();
                    return;
                }

                this.loadAnalytics(contractorName);
            });
        }

        const objectButtons = this.container.querySelectorAll('[data-ca-object]');
        objectButtons.forEach((button) => {
            button.addEventListener('click', () => {
                const customerName = button.getAttribute('data-ca-customer') || '';
                const objectName = button.getAttribute('data-ca-object') || '';
                this.loadObjectDetails(customerName, objectName);
            });
        });
    }

    renderSummary() {
        const summary = this.analytics?.summary;
        if (!summary) return '';

        return `
            <div class="ca-summary-grid">
                <div class="ca-summary-card">
                    <div class="ca-summary-label">Цена контрактов</div>
                    <div class="ca-summary-value">${formatMoney(summary.contracts_sum)}</div>
                </div>
                <div class="ca-summary-card">
                    <div class="ca-summary-label">Объектов</div>
                    <div class="ca-summary-value">${summary.objects_count ?? 0}</div>
                </div>
                <div class="ca-summary-card">
                    <div class="ca-summary-label">Просроченные объекты</div>
                    <div class="ca-summary-value">${summary.overdue_objects_count ?? 0}</div>
                </div>
            </div>
        `;
    }

    renderTree() {
        const customers = this.analytics?.customers || [];
        const selected = this.analytics?.selected_object;
        const selectedKey = `${selected?.customer_name || ''}::${selected?.object_name || ''}`;

        if (customers.length === 0) {
            return '<div class="ca-empty">Нет данных по структуре заказчиков и объектов</div>';
        }

        return `
            <div class="ca-tree-root-node">${escapeHtml(this.analytics.contractor_name || this.selectedContractor)}</div>
            <div class="ca-customers">
                ${customers.map((customer) => `
                    <section class="ca-customer-group" data-ca-group>
                        <svg class="ca-links-layer" data-ca-links aria-hidden="true"></svg>
                        <div class="ca-customer-node" data-ca-customer-node>
                            <div class="ca-customer-name">${escapeHtml(customer.customer_name || 'Без названия')}</div>
                            <div class="ca-customer-meta">${customer.objects_count || 0} объекта</div>
                        </div>
                        <div class="ca-objects-list">
                            ${(customer.objects || []).map((object) => {
        const key = `${object.customer_name || customer.customer_name}::${object.object_name || ''}`;
        const selectedClass = key === selectedKey ? ' ca-object-selected' : '';
        const overdueAmount = Number(object.overdue_debt_amount);
        const isOverdueObject = Number.isFinite(overdueAmount) && overdueAmount > 0;
        const overdueClass = isOverdueObject ? ' ca-object-overdue' : '';

        return `
                                    <div class="ca-object-row">
                                        <button
                                            type="button"
                                            class="ca-object-node${selectedClass}${overdueClass}"
                                            data-ca-customer="${escapeHtml(object.customer_name || customer.customer_name || '')}"
                                            data-ca-object="${escapeHtml(object.object_name || '')}"
                                            data-ca-object-node
                                        >
                                            <span class="ca-object-dot ${riskClass(object.risk_level)}"></span>
                                            <span class="ca-object-content">
                                                <span class="ca-object-name">${escapeHtml(object.object_name || 'Без названия')}</span>
                                                <span class="ca-object-meta">
                                                    ${formatMoney(object.debet_sum)}
                                                </span>
                                            </span>
                                        </button>
                                    </div>
                                `;
    }).join('')}
                        </div>
                    </section>
                `).join('')}
            </div>
        `;
    }

    renderDetails() {
        const details = this.analytics?.selected_object;
        if (!details) {
            return '<div class="ca-empty">Выберите объект для просмотра деталей</div>';
        }

        const status = statusMeta(details);
        const readinessText = formatReadiness(details.readiness_percent);
        const readinessPercent = parseReadinessToPercent(details.readiness_percent);
        const progress = readinessPercent === null ? 0 : readinessPercent;

        return `
            <div class="ca-details">
                <div class="ca-details-header">
                    <h3>${escapeHtml(details.object_name || 'Объект')}</h3>
                    <span class="ca-status-badge ${status.className}">${status.text}</span>
                </div>

                <div class="ca-details-grid">
                    <div class="ca-detail-row"><span>Заказчик</span><strong>${escapeHtml(details.customer_name || '—')}</strong></div>
                    <div class="ca-detail-row"><span>Подрядчик</span><strong>${escapeHtml(details.contractor_name || '—')}</strong></div>
                    <div class="ca-detail-row"><span>Цена контракта</span><strong>${formatMoney(details.contract_sum)}</strong></div>
                    <div class="ca-detail-row"><span>Кассовые расходы</span><strong>${formatMoney(details.paid_sum)}</strong></div>
                    <div class="ca-detail-row"><span>Процент готовности</span><strong>${escapeHtml(readinessText)}</strong></div>
                    <div class="ca-detail-row"><span>ТДЦ</span><strong>${formatMoney(details.tdc_sum)}</strong></div>
                    <div class="ca-detail-row"><span>РВ</span><strong>${details.rv_exists ? 'Да' : 'Нет'}</strong></div>
                    <div class="ca-detail-row"><span>Дебиторская задолженность</span><strong>${formatMoney(details.debet_sum)}</strong></div>
                    <div class="ca-detail-row"><span>Просроченная задолженность</span><strong class="ca-overdue-value">${formatMoney(details.overdue_debt_amount)}</strong></div>
                    <div class="ca-detail-row"><span>Объём принятых работ</span><strong>${formatPercent(details.accepted_percent, 1)}</strong></div>
                </div>

                <div class="ca-progress-section">
                    <div class="ca-progress-header">
                        <span>Прогресс работ</span>
                        <strong>${escapeHtml(readinessText)}</strong>
                    </div>
                    <div class="ca-progress-track">
                        <div class="ca-progress-fill" style="width:${progress}%;"></div>
                    </div>
                    <div class="ca-progress-dates">
                        <div>
                            <div class="ca-progress-date">${formatDate(details.work_start_date)}</div>
                            <div class="ca-progress-label">Начало работ</div>
                        </div>
                        <div>
                            <div class="ca-progress-date">${formatDate(details.work_end_date)}</div>
                            <div class="ca-progress-label">Окончание работ</div>
                        </div>
                    </div>
                </div>
            </div>
        `;
    }

    render() {
        const contractorsOptions = this.contractors
            .map((name) => `<option value="${escapeHtml(name)}" ${name === this.selectedContractor ? 'selected' : ''}>${escapeHtml(name)}</option>`)
            .join('');

        this.container.innerHTML = `
            <div class="contractor-analytics-root">
                <div class="ca-toolbar">
                    <label for="ca-select">Подрядчик</label>
                    <select id="ca-select" data-ca-select ${this.loadingContractors ? 'disabled' : ''}>
                        <option value="">-- Выберите подрядчика --</option>
                        ${contractorsOptions}
                    </select>
                </div>

                ${this.error ? `<div class="ca-error">${escapeHtml(this.error)}</div>` : ''}
                ${this.loadingContractors ? '<div class="ca-loading">Загрузка списка подрядчиков...</div>' : ''}
                ${this.loading && this.analytics ? '<div class="ca-loading-inline">Обновление данных...</div>' : ''}
                ${this.emptyStateMessage ? `
                    <div class="ca-empty-state">
                        <div class="ca-empty-state-title">Нет данных для отображения</div>
                        <div class="ca-empty-state-text">${escapeHtml(this.emptyStateMessage)}</div>
                    </div>
                ` : ''}
                ${this.analytics ? `
                    ${this.renderSummary()}
                    <div class="ca-content">
                        <div class="ca-tree-pane">${this.renderTree()}</div>
                        <div class="ca-details-pane">${this.renderDetails()}</div>
                    </div>
                    <div class="ca-legend">
                        <span><i class="text"></i>Степень готовности объекта: </span>
                        <span><i class="ca-object-dot ca-risk-ok"></i>(>=70%)</span>
                        <span><i class="ca-object-dot ca-risk-risk"></i>(30-70%)</span>
                        <span><i class="ca-object-dot ca-risk-critical"></i>(<30%)</span>
                        <span><i class="ca-object-dot ca-risk-no-data"></i>Нет данных</span>
                    </div>
                    <div class="ca-legend">
                        <span><i class="text"></i>Тип дебиторской задолженности: </span>
                        <span><i class="ca-object-dot ca-risk-ok"></i>Просроченная дебиторская задолженность</span>
                        <span><i class="ca-object-dot ca-risk-risk"></i>Текущая дебиторская задоженность</span>
                    </div>    
                    
                ` : ''}
            </div>
        `;

        this.bindEvents();
        this.scheduleConnectionsRender();
    }
}
