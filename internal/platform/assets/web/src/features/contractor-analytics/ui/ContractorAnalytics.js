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
    return new Intl.NumberFormat('ru-RU', {
        maximumFractionDigits: 0
    }).format(amount) + ' ₽';
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
    const raw = String(details?.status || '').toLowerCase();
    if (raw.includes('проср')) return { text: 'Просрочен', className: 'ca-status-overdue' };
    if (raw.includes('работ')) return { text: 'В работе', className: 'ca-status-active' };
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

            // Детальную карточку справа показываем только после явного клика по объекту.
            this.analytics = {
                ...data,
                selected_object: null
            };
            this.emptyStateMessage = '';
        } catch (err) {
            if (token !== this.requestToken) return;
            const message = String(err?.message || '');
            this.emptyStateMessage = '';
            if (message.includes('contractor has no objects')) {
                this.emptyStateMessage = 'По выбранному подрядчику нет объектов (construction_object).';
                this.error = null;
                this.analytics = null;
                return;
            }
            this.error = `Не удалось загрузить аналитику: ${message}`;
            if (false) {
                this.error = 'По выбранному подрядчику нет объектов (construction_object).';
            } else {
                this.error = `Не удалось загрузить аналитику: ${message}`;
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
                    <div class="ca-summary-label">Сумма контрактов</div>
                    <div class="ca-summary-value">${formatMoney(summary.contracts_sum)}</div>
                </div>
                <div class="ca-summary-card">
                    <div class="ca-summary-label">Объектов</div>
                    <div class="ca-summary-value">${summary.objects_count ?? 0}</div>
                </div>
                <div class="ca-summary-card">
                    <div class="ca-summary-label">Средняя готовность</div>
                    <div class="ca-summary-value">${formatPercent(summary.avg_readiness_percent)}</div>
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
                    <section class="ca-customer-group">
                        <div class="ca-customer-node">
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
        const readiness = object.readiness_percent === null || object.readiness_percent === undefined
            ? 'нет данных'
            : formatPercent(object.readiness_percent);

        return `
                                    <div class="ca-object-row">
                                        <span class="ca-link-connector" aria-hidden="true"></span>
                                        <button
                                            type="button"
                                            class="ca-object-node${selectedClass}${overdueClass}"
                                            data-ca-customer="${escapeHtml(object.customer_name || customer.customer_name || '')}"
                                            data-ca-object="${escapeHtml(object.object_name || '')}"
                                        >
                                            <span class="ca-object-dot ${riskClass(object.risk_level)}"></span>
                                            <span class="ca-object-content">
                                                <span class="ca-object-name">${escapeHtml(object.object_name || 'Без названия')}</span>
                                                <span class="ca-object-meta">
                                                    ${formatMoney(object.contract_sum)} · ${escapeHtml(readiness)}${isOverdueObject ? ' · Просрочка' : ''}
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
        const readiness = Number(details.readiness_percent);
        const progress = Number.isNaN(readiness) ? 0 : Math.max(0, Math.min(100, readiness));

        return `
            <div class="ca-details">
                <div class="ca-details-header">
                    <h3>${escapeHtml(details.object_name || 'Объект')}</h3>
                    <span class="ca-status-badge ${status.className}">${status.text}</span>
                </div>

                <div class="ca-details-grid">
                    <div class="ca-detail-row"><span>Заказчик</span><strong>${escapeHtml(details.customer_name || '—')}</strong></div>
                    <div class="ca-detail-row"><span>Подрядчик</span><strong>${escapeHtml(details.contractor_name || '—')}</strong></div>
                    <div class="ca-detail-row"><span>Сумма контракта</span><strong>${formatMoney(details.contract_sum)}</strong></div>
                    <div class="ca-detail-row"><span>Перечислено</span><strong>${formatMoney(details.paid_sum)}</strong></div>
                    <div class="ca-detail-row"><span>Процент готовности</span><strong>${formatPercent(details.readiness_percent)}</strong></div>
                    <div class="ca-detail-row"><span>ТДЦ</span><strong>${formatMoney(details.tdc_sum)}</strong></div>
                    <div class="ca-detail-row"><span>РВ</span><strong>${details.rv_exists ? 'Да' : 'Нет'}</strong></div>
                    <div class="ca-detail-row"><span>Дебиторская задолженность</span><strong>${formatMoney(details.debet_sum)}</strong></div>
                    <div class="ca-detail-row"><span>Просроченная задолженность</span><strong class="ca-overdue-value">${formatMoney(details.overdue_debt_amount)}</strong></div>
                    <div class="ca-detail-row"><span>Принято</span><strong>${formatPercent(details.accepted_percent, 1)}</strong></div>
                </div>

                <div class="ca-progress-section">
                    <div class="ca-progress-header">
                        <span>Прогресс работ</span>
                        <strong>${formatPercent(details.readiness_percent)}</strong>
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
                        <span><i class="ca-object-dot ca-risk-ok"></i>(>=70%)</span>
                        <span><i class="ca-object-dot ca-risk-risk"></i>(30-70%)</span>
                        <span><i class="ca-object-dot ca-risk-critical"></i>(<30%)</span>
                        <span><i class="ca-object-dot ca-risk-no-data"></i> Нет данных</span>
                    </div>
                ` : ''}
            </div>
        `;

        this.bindEvents();
    }
}
