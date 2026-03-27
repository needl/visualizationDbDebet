// src/components/PieChartComponent.js
import { appState } from '../state/appState.js';

export class PieChartComponent {
    constructor(container, metricKey, title) {
        this.container = container;
        this.metricKey = metricKey;
        this.title = title;
        this.chart = null;

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
            console.error('ECharts не загружен');
            this.showError('ECharts не загружен');
            return;
        }

        if (!this.chart) {
            this.chart = echarts.init(this.container);
        }

        const option = {
            tooltip: {
                trigger: 'item',
                position: function(point, params, dom, rect, size) {
                    return [size.viewSize[0] / 2, size.viewSize[1] / 2];
                },
                formatter: function(params) {
                    const valueInBillions = (params.value / 1_000_000_000).toLocaleString('ru-RU', { maximumFractionDigits: 2 });
                    return `${params.name}: ${valueInBillions} млрд ₽ (${params.percent}%)`;
                }
            },
            legend: {
                orient: 'vertical',
                left: 'left',
                top: 'middle',
                itemWidth: 20,
                itemHeight: 14,
                textStyle: { fontSize: 11 }
            },
            series: [
                {
                    name: this.title,
                    type: 'pie',
                    radius: ['45%', '65%'],
                    center: ['50%', '50%'],
                    avoidLabelOverlap: false,
                    itemStyle: {
                        borderRadius: 8,
                        borderColor: '#fff',
                        borderWidth: 2
                    },
                    label: { show: false },
                    emphasis: {
                        label: { show: true, fontSize: 20, fontWeight: 'bold' }
                    },
                    labelLine: { show: false },
                    data: data
                }
            ]
        };
        this.chart.setOption(option);
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
}