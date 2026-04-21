// src/components/PieChartComponent.js
import { appState } from '../state/appState.js';

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

        // Очищаем контейнер от всего (сообщение, предыдущий canvas)
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
        this.container.innerHTML = '';

        // Инициализируем новый экземпляр ECharts
        this.chart = echarts.init(this.container);

        const option = {
            /*tooltip: {
                trigger: 'item',
                position: function(point, param33s, dom, rect, size) {
                    // Показываем тултип рядом с курсором
                    return [point[0] / 2, point[1] / 1];
                },
                formatter: function(params) {
                    const valueInBillions = (params.value / 1_000_000_000).toLocaleString('ru-RU', { maximumFractionDigits: 2 });
                    return `${valueInBillions} млрд ₽ (${params.percent}%)`;
                }
            },*/
            tooltip: {
                trigger: 'item',
                confine: true,
                appendTo: document.body,
                // position можно не указывать, ECharts сам подберёт
                formatter: function(params) {
                    const valueInBillions = (params.value / 1_000_000_000).toLocaleString('ru-RU', { maximumFractionDigits: 2 });
                    return `${valueInBillions} млрд ₽ (${params.percent}%)`;
                }
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
                    width: 200,            // максимальная ширина текста в пикселях
                    overflow: 'break'      // или 'truncate' с многоточием
                },
                formatter: function (name) {
                    // Опционально: ручной перенос длинных слов
                    if (name.length > 20) {
                        return name.slice(0, 20) + '…';
                    }
                    return name;
                }
            },
            /*legend: {
                orient: 'vertical',
                left: 50,
                top: 'middle',
                itemWidth: 50,
                itemHeight: 20,
                textStyle: { fontSize: 14 }
            },*/
            series: [
                {
                    name: this.title,
                    type: 'pie',
                    radius: ['5%', '90%'],
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
                            fontSize: 20,
                            position: 'top',
                            fontWeight: 'bold',
                        }
                    },
                    labelLine: { show: false },
                    data: data
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