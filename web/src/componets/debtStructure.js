import { appState } from '../state/appState.js';

export class DebtStructure {
    constructor(container) {
        this.container = container;
        this.chart = null;
    }

    render(summary) {

        if (!summary || summary.total_debet === undefined) {
            this.container.innerHTML = '<div class="empty-message">Нет данных по задолженности</div>';
            return;
        }

        const totalDebt = summary.total_debet;
        const overdueDebt = summary.total_debet_overdue || 0;
        const longTermDebt = summary.total_debet_long || 0;
        const currentDebt = totalDebt - overdueDebt - longTermDebt;

        if (totalDebt === 0 && overdueDebt === 0 && longTermDebt === 0) {
            this.container.innerHTML = '<div class="empty-message">Задолженность отсутствует</div>';
            return;
        }

        if (!this.chart) {
            this.chart = echarts.init(this.container);
        }

        // Подготавливаем данные, исключая нулевые значения
        const data = [
            { name: 'Текущая задолженность', value: currentDebt, itemStyle: { color: '#10b981' } },
            { name: 'Просроченная задолженность', value: overdueDebt, itemStyle: { color: '#ef4444' } },
            { name: 'Долгосрочная задолженность', value: longTermDebt, itemStyle: { color: '#f59e0b' } }
        ].filter(item => item.value > 0);

        if (data.length === 0) {
            this.container.innerHTML = '<div class="empty-message">Задолженность отсутствует</div>';
            return;
        }

        const option = {
            title: {
                text: 'Структура дебиторской задолженности',
                left: 'center',
                top: 0
            },
            tooltip: {
                trigger: 'item',
                formatter: (params) => {
                    const value = params.value;
                    const formatted = new Intl.NumberFormat('ru-RU', {
                        style: 'currency',
                        currency: 'RUB',
                        minimumFractionDigits: 0,
                        maximumFractionDigits: 0
                    }).format(value);
                    return `${params.name}: ${formatted} (${params.percent}%)`;
                }
            },
            series: [
                {
                    name: 'Дебиторская задолженность',
                    type: 'pie',
                    radius: '75%',
                    avoidLabelOverlap: false,
                    label: {
                        show: true,
                        position: 'outside',
                        formatter: (params) => {
                            const percent = params.percent.toFixed(1);
                            return `${params.name}\n${percent}%`;
                        }
                    },
                    emphasis: {
                        scale: true
                    },
                    data: data
                }
            ]
        };

        this.chart.setOption(option, true);
        this.chart.resize();
    }

    clear() {
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
        this.container.innerHTML = '';
    }
}