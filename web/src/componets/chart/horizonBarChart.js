// src/components/horizonBarChart.js
export class HorizontalBarChart {
    constructor(container, title, valueFormatter, onBarClick) {
        this.container = container;
        this.title = title;
        this.valueFormatter = valueFormatter || ((v) => (v / 1e9).toFixed(2) + ' млрд ₽');
        this.onBarClick = onBarClick || null;   // <-- сохраняем
        this.chart = null;
    }

    render(data) {
        if (!data || data.length === 0) {
            this.container.innerHTML = '<div class="empty-message">Нет данных по топ-10</div>';
            if (this.chart) {
                this.chart.dispose();
                this.chart = null;
            }
            return;
        }

        if (!this.chart) {
            this.chart = echarts.init(this.container);
        }

        const names = data.map(item => item.name);
        const values = data.map(item => item.value);

        const option = {
            title: {
                text: this.title,
                show: false,
                left: 'center',
                textStyle: { fontSize: 14, fontWeight: 'bold' }
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: { type: 'shadow' },
                confine: true,
                formatter: (params) => `${params[0].name}`
            },
            grid: {
                containLabel: true,
                left: '3%',
                right: '7%',
                top: '5%',
                bottom: '5%'
            },
            xAxis: {
                type: 'value',
                name: 'Сумма (млрд ₽)',
                nameLocation: 'middle',
                nameGap: 28,
                axisLabel: {
                    formatter: (value) => (value / 1e9).toFixed(1) + ' млрд'
                }
            },
            yAxis: {
                type: 'category',
                data: names,
                inverse: true,
                axisLabel: {
                    width: 300,
                    overflow: 'break',
                    fontSize: 10,
                    rotate: 0,
                    interval: 0
                }
            },
            series: [{
                name: this.title,
                type: 'bar',
                data: values,
                cursor: 'pointer',                // визуально указываем, что можно кликать
                itemStyle: { color: '#d11b1b' },
                label: {
                    show: true,
                    position: 'right',
                    formatter: (p) => this.valueFormatter(p.value)
                }
            }]
        };

        // Снимаем старый обработчик и задаём новый
        this.chart.off('click');
        this.chart.setOption(option, true);

        if (this.onBarClick) {
            this.chart.on('click', (params) => {
                if (params.componentType === 'series' && params.dataIndex !== undefined) {
                    const name = names[params.dataIndex];
                    const value = values[params.dataIndex];
                    this.onBarClick(name, value);
                }
            });
        }

        this.chart.resize();
    }

    resize() {
        if (this.chart) this.chart.resize();
    }
}