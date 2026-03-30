export class DebtStructure {
    constructor(container) {
        this.container = container;
        this.chart = null;
    }

    render(summary) {
        this.clear();

        if (!summary || summary.total_debet === undefined) {
            this.container.innerHTML = '<div class="empty-message">Нет данных по задолженности</div>';
            return;
        }

        const totalDebt = summary.total_debet;
        const overdueDebt = summary.total_debet_overdue || 0;
        const currentDebt = totalDebt - overdueDebt;

        if (totalDebt === 0 && overdueDebt === 0) {
            this.container.innerHTML = '<div class="empty-message">Нет информации по задолженности</div>';
            return;
        }

        if (!this.chart) {
            this.chart = echarts.init(this.container);
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
                    radius: '60%',           // полноценный круг (без внутреннего кольца)
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
                    data: [
                        { name: 'Текущая задолженность', value: currentDebt, itemStyle: { color: '#d17e2d' } },
                        { name: 'Просроченная задолженность', value: overdueDebt, itemStyle: { color: '#d11b1b' } }
                    ]
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