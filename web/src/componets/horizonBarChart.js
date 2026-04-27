export class HorizontalBarChart {
    constructor(container, title, valueFormatter = (v) => (v / 1e9).toFixed(2) + ' млрд ₽') {
        this.container = container;
        this.title = title;
        this.valueFormatter = valueFormatter;
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
                textStyle: {
                    fontSize: 14,
                    fontWeight: 'bold'
                }
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: { type: 'shadow' },
                confine: true,
                position: 'centre',
                formatter: (params) => {
                    const val = params[0].value;
                    //return `${params[0].name}<br/>${this.valueFormatter(val)}`;
                    return `${params[0].name}`;
                }
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
                axisLabel: {
                    width: 300,               // максимальная ширина метки в px
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
                itemStyle: { color: '#d11b1b' },
                label: {
                    show: true,
                    position: 'right',
                    formatter: (p) => this.valueFormatter(p.value)
                }
            }]
        };

        this.chart.setOption(option, true);
        this.chart.resize();
    }

    resize() {
        if (this.chart) this.chart.resize();
    }
}