/* global echarts */

import { appState } from '../state/appState.js';

export class ChartComponent {
    constructor(container, metricKey, title) {
        this.container = container; // это уже DOM-элемент, а не id
        this.metricKey = metricKey;
        this.title = title;
        this.chart = null;
        this.resizeHandler = null;

        // Подписка на изменение размера окна
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
        // Определяем тип данных: если есть series — групповой, иначе обычный
        if (chartData.series && Array.isArray(chartData.series)) {
            // Групповой график
            this.drawGroupedChart(chartData.names, chartData.series);
        } /*else if (chartData.values) {
            // Обычный график
            this.drawChart(chartData.names, chartData.values);
        }*/ else {
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

        // Цвета для четырёх серий
        const colorMap = {
            '31.12.2024 Текущая дебиторская задолженность': '#4caf50',
            '31.12.2024 Просроченная дебиторская задолженность': '#f44336',
            '31.12.2025 Текущая дебиторская задолженность': '#81c784',
            '31.12.2025 Просроченная дебиторская задолженность': '#e57373'
        };

        const option = {
            /*legend: {
                bottom: 40,
                left: 'center',
                orient: 'horizontal',
                itemGap: 10
            },*/
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
            /*tooltip: {
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
            },*/
            tooltip: {
                trigger: 'item',   // было 'axis'
                formatter: function(params) {
                    const val = params.value;
                    const inMillions = (val / 1_000_000_000).toLocaleString('ru-RU', { maximumFractionDigits: 2 });
                    return `${params.seriesName}: ${inMillions} млрд ₽`;
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
            series: series.map(s => ({
                name: s.name,
                type: 'bar',
                stack: s.stack,               // <-- главное изменение
                data: s.data,
                itemStyle: {
                    color: colorMap[s.name] || undefined
                }
            }))
        };

        this.chart.setOption(option);
        this.chart.resize();
    }

    /*drawGroupedChart(names, series) {
        if (!this.container) return;
        if (typeof echarts === 'undefined') {
            this.showError('ECharts не загружен');
            return;
        }

        if (!this.chart) {
            this.chart = echarts.init(this.container);
        }

        const option = {
            legend: {
                bottom: 40,
                left: 'center',
                orient: 'horizontal',
                itemGap: 10
            },
            title: {
                text: this.title,
                left: 'center'
            },
            grid: {
                containLabel: true,
                left: '2%',
                right: '2%',
                top: '10%',
                bottom: '10%'
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: { type: 'shadow' },
                position: 'auto', // автоматическое позиционирование
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
                        return (value / 1_000_000_000).toFixed(1);
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
        this.chart.resize(); // обновляем размеры
    }*/

    /*drawChart(names, values) {
        if (!this.container) return;
        if (typeof echarts === 'undefined') {
            this.showError('ECharts не загружен');
            return;
        }

        if (!this.chart) {
            this.chart = echarts.init(this.container);
        }

        const option = {
            title: {
                text: this.title,
                left: 'center'
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: { type: 'shadow' },
                // Улучшенное позиционирование: автоматически подстраивается
                position: 'auto',
                formatter: function(params) {
                    const data = params[0];
                    const roundedValue = Math.round(data.value);
                    const inMillions = (roundedValue / 1000000000).toLocaleString('ru-RU', { maximumFractionDigits: 2 });
                    return `${data.name}<br/>${data.seriesName}: ${inMillions} млрд ₽`;
                }
            },
            grid: {
                containLabel: true,
                left: '8%',
                right: '5%',
                top: '15%',
                bottom: '10%'
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
                    color: '#111ecd'
                }
            }]
        };
        this.chart.setOption(option);
        this.chart.resize(); // принудительно обновляем размер
    }*/

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
        // Удаляем обработчик resize
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