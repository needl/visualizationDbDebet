import { appState } from '../../../shared/state/appState.js';

export class PieChartComponent {
    constructor(container, metricKey, title) {
        this.container = container;
        this.metricKey = metricKey;
        this.title = title;
        this.chart = null;
        this.resizeHandler = null;

        this.resizeHandler = () => {
            if (this.chart) this.chart.resize();
        };
        window.addEventListener('resize', this.resizeHandler);

        appState.subscribe((state) => {
            this.render(state);
        });
    }

    render(state) {
        if (state.loading) {
            this.showLoading();
            return;
        }
        if (state.error) {
            this.showError(state.error);
            return;
        }
        const chartData = state.chartData[this.metricKey];
        if (!chartData || !chartData.data || chartData.data.length === 0) {
            this.showEmpty();
            return;
        }
        this.drawPie(chartData.data);
    }

    drawPie(data) {
        if (!this.container) return;
        if (typeof echarts === 'undefined') {
            this.showError('ECharts не загружен');
            return;
        }

        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
        this.container.innerHTML = '';
        this.chart = echarts.init(this.container);

        const formatValue = (value, percent) => {
            const valueInBillions = (value / 1_000_000_000).toLocaleString('ru-RU', { maximumFractionDigits: 2 });
            return `${valueInBillions} млрд ₽ (${percent}%)`;
        };

        const tooltipFormatter = (params) => formatValue(params.value, params.percent);

        const legendTooltipFormatter = (params) => {
            const item = data.find((d) => d.name === params.name);
            if (item) {
                const total = data.reduce((sum, d) => sum + d.value, 0);
                const percent = ((item.value / total) * 100).toFixed(2);
                return formatValue(item.value, percent);
            }
            return params.name;
        };

        const option = {
            tooltip: {
                trigger: 'item',
                confine: true,
                appendTo: document.body,
                formatter: tooltipFormatter
            },
            legend: {
                orient: 'vertical',
                left: 50,
                top: 'middle',
                itemWidth: 20,
                itemHeight: 12,
                itemGap: 12,
                textStyle: {
                    fontSize: 14,
                    width: 200,
                    overflow: 'break'
                },
                formatter: function(name) {
                    if (name.length > 20) {
                        return `${name.slice(0, 20)}…`;
                    }
                    return name;
                },
                tooltip: {
                    show: true,
                    formatter: legendTooltipFormatter
                }
            },
            series: [
                {
                    name: this.title,
                    type: 'pie',
                    radius: ['5%', '85%'],
                    center: ['63%', '50%'],
                    avoidLabelOverlap: false,
                    itemStyle: {
                        borderRadius: 8,
                        borderColor: '#ffffff',
                        borderWidth: 2
                    },
                    label: { show: false },
                    emphasis: {
                        scale: false,
                        label: {
                            show: true,
                            fontSize: 12,
                            position: 'outer',
                            fontWeight: 'bold'
                        }
                    },
                    labelLine: { show: false },
                    data
                }
            ]
        };

        this.chart.setOption(option, true);
        this.chart.resize();
    }

    showLoading() {
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
        this.container.innerHTML = '';
        const loadingDiv = document.createElement('div');
        loadingDiv.className = 'loading';
        loadingDiv.textContent = 'Загрузка данных...';
        this.container.appendChild(loadingDiv);
    }

    showError(errorMsg) {
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
        this.container.innerHTML = '';
        const errorDiv = document.createElement('div');
        errorDiv.className = 'error';
        errorDiv.textContent = `Ошибка: ${errorMsg}`;
        this.container.appendChild(errorDiv);
    }

    showEmpty() {
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
        this.container.innerHTML = '';
        const emptyDiv = document.createElement('div');
        emptyDiv.className = 'empty';
        emptyDiv.textContent = 'Нет данных';
        this.container.appendChild(emptyDiv);
    }

    dispose() {
        if (this.resizeHandler) {
            window.removeEventListener('resize', this.resizeHandler);
            this.resizeHandler = null;
        }
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
        this.container.innerHTML = '';
    }
}
