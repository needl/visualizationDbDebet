

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
        const currentDebt = totalDebt - overdueDebt;

        if (totalDebt === 0 && overdueDebt === 0 && longTermDebt === 0) {
            this.container.innerHTML = '<div class="empty-message">Задолженность отсутствует</div>';
            return;
        }

        if (!this.chart) {
            this.chart = echarts.init(this.container);
        }

        // Подготавливаем данные, исключая нулевые значения
        const data = [
            { name: 'Текущая', value: currentDebt, itemStyle: { color: '#10b981' } },
            { name: 'Просроченная', value: overdueDebt, itemStyle: { color: '#ef4444' } },
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
                    radius: '80%',
                    avoidLabelOverlap: false,
                    label: {
                        show: true,
                        position: 'outside',
                        formatter: (params) => {
                            const percent = params.percent.toFixed(1);
                            const nameLines = params.name.split(' ').join('\n');
                            return `${nameLines}\n${percent}%`;
                        }
                    },
                    labelLayout: (params) => {
                        // Определяем, какую метку двигаем влево, а какую вправо
                        // Можно ориентироваться на название или индекс
                        if (params.text.includes('Просроченная')) {
                            // Смещаем левее и выравниваем по правому краю
                            return {
                                x: params.labelRect.x,  // сдвиг влево (подбирается визуально)
                                y: params.labelRect.y + 30,
                                align: 'centre'
                            };
                        } else {
                            // Смещаем правее и выравниваем по левому краю
                            return {
                                x: params.labelRect.x,  // сдвиг вправо
                                y: params.labelRect.y,
                                align: 'centre'
                            };
                        }
                    },
                    labelLine: {
                        length: 15,      // длина первой части линии от сектора
                        length2: 10,      // длина второй части (горизонтальной)
                        smooth: true      // плавное соединение
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