// src/components/CustomerCard.js
import {appState} from "../state/appState.js";

export class CustomerCard {
    constructor(container) {
        this.container = container;
        this.unsubscribe = null;
    }

    render(summary) {
        if (!summary) {
            this.container.innerHTML = '<div class="loading">Загрузка статистики...</div>';
            return;
        }
        const formatCurrency = (val) => (val / 1e9).toFixed(2) + ' млрд ₽';
        const formatPercent = (val) => val.toFixed(2) + '%';
        this.container.innerHTML = `
            <div class="stats-grid" style="margin-top:0">
                <div class="metric-card-wrapper"><div class="metric-card"><div class="card-title">Кол-во контрагентов</div><div class="card-value">${summary.contractors_count}</div></div></div>
                <div class="metric-card-wrapper"><div class="metric-card"><div class="card-title">Сумма контрактов</div><div class="card-value">${formatCurrency(summary.total_contract_amount)}</div></div></div>
                <div class="metric-card-wrapper"><div class="metric-card"><div class="card-title">Перечислено</div><div class="card-value">${formatCurrency(summary.total_paid_amount)}</div></div></div>
                <div class="metric-card-wrapper"><div class="metric-card"><div class="card-title">Принято работ</div><div class="card-value">${formatCurrency(summary.total_accepted_amount)}</div></div></div>
                <div class="metric-card-wrapper"><div class="metric-card"><div class="card-title">% принятых</div><div class="card-value">${formatPercent(summary.percentage)}</div></div></div>
            </div>
        `;
    }

    mount() {
        this.unsubscribe = appState.subscribe(state => {
            if (state.customerLoading) {
                this.container.innerHTML = '<div class="loading">Загрузка статистики...</div>';
            } else if (state.customerSummary) {
                this.render(state.customerSummary);
            } else {
                this.container.innerHTML = '<div class="empty-message">Выберите заказчика</div>';
            }
        });
    }

    unmount() {
        if (this.unsubscribe) this.unsubscribe();
    }
}