import { appState } from '../../../shared/state/appState.js';

export class ChartComponent {
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
        if (!chartData || !chartData.names) {
            this.showEmpty();
            return;
        }
        if (chartData.series && Array.isArray(chartData.series)) {
            this.drawGroupedChart(chartData.names, chartData.series);
        } else {
            this.showEmpty();
        }
    }

    drawGroupedChart(names, series) {
        if (!this.container) return;
        if (typeof echarts === 'undefined') {
            this.showError('ECharts не загружен');
            return;
        }

        if (!this.chart) {
            this.chart = echarts.init(this.container);
        }

        const colorMap = {
            '31.12.2024 Текущая дебиторская задолженность': '#81c784',
            '31.12.2024 Просроченная дебиторская задолженность': '#e57373',
            '31.12.2025 Текущая дебиторская задолженность': '#4caf50',
            '31.12.2025 Просроченная дебиторская задолженность': '#f44336'
        };

        const option = {
            title: {
                text: this.title,
                left: 'center'
            },
            grid: {
                containLabel: true,
                left: '2%',
                right: '2%',
                top: '10%',
                bottom: '2%'
            },
            tooltip: {
                trigger: 'item',
                formatter: function(params) {
                    const val = params.value;
                    const inBillions = (val / 1_000_000_000).toLocaleString('ru-RU', { maximumFractionDigits: 2 });
                    return `${params.seriesName}: ${inBillions} млрд ₽`;
                }
            },
            xAxis: {
                type: 'category',
                data: names,
                axisLabel: {
                    rotate: 45,
                    interval: 0
                }
            },
            yAxis: {
                type: 'value',
                name: 'Сумма (млрд ₽)',
                axisLabel: {
                    formatter: function(value) {
                        return (value / 1_000_000_000).toFixed(1);
                    }
                }
            },
            series: series.map((s) => ({
                name: s.name,
                type: 'bar',
                stack: s.stack,
                data: s.data,
                itemStyle: {
                    color: colorMap[s.name] || undefined
                }
            }))
        };

        this.chart.setOption(option);
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
