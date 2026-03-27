/* global echarts */

import { appState } from '../state/appState.js';

export class ChartComponent {
    constructor(container, metricKey, title) {
        this.container = container; // это уже DOM-элемент, а не id
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
        if (!chartData || !chartData.names) {
            this.showEmpty();
            return;
        }
        // Определяем тип данных: если есть series — групповой, иначе обычный
        if (chartData.series && Array.isArray(chartData.series)) {
            // Групповой график
            this.drawGroupedChart(chartData.names, chartData.series);
        } else if (chartData.values) {
            // Обычный график
            this.drawChart(chartData.names, chartData.values);
        } else {
            this.showEmpty();
        }
    }

    drawChart(names, values) {
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
            // legend: { bottom: 0, left: 'center', orient: 'horizontal' },
            title: {
                text: this.title,
                left: 'center'
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: { type: 'shadow' },
                position: function(point, params, dom, rect, size) {
                    // size.viewSize — массив [ширина, высота] области графика
                    return [size.viewSize[0] / 6.5, size.viewSize[1] /3];
                },
                formatter: function(params) {
                    const data = params[0];
                    const roundedValue = Math.round(data.value);
                    const inMillions = (roundedValue / 1000000000).toLocaleString('ru-RU', { maximumFractionDigits: 2 });
                    return `${data.name}<br/>${data.seriesName}: ${inMillions} млрд ₽`;
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
                        return (value / 1000000000).toFixed(0);
                    }
                }
            },
            series: [{
                name: this.title,
                type: 'bar',
                data: values,
                itemStyle: {
                    color: '#cd1111'
                }
            }]
        };
        this.chart.setOption(option);
    }

    drawGroupedChart(names, series) {
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
            legend: { bottom: 0, left: 'center', orient: 'horizontal' },
            title: {
                text: this.title,
                left: 'center'
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: { type: 'shadow' },
                formatter: function(params) {
                    let res = params[0].axisValue + '<br/>';
                    params.forEach(p => {
                        const val = p.value;
                        const inMillions = (val / 1_000_000_000).toLocaleString('ru-RU', { maximumFractionDigits: 2 });
                        res += `${p.seriesName}: ${inMillions} млрд ₽<br/>`;
                    });
                    return res;
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
                        return (value / 1_000_000_000).toFixed(0);
                    }
                }
            },
            series: series.map(s => ({
                name: s.name,
                type: 'bar',
                data: s.data,
                itemStyle: { color: s.name === '2024' ? '#da8608' : '#cd1111' }
            }))
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