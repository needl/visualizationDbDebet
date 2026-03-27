// src/dashboardRenderer.js
import { ChartComponent } from "./componets/chartForDeb.js";
import {PieChartComponent} from "./componets/pieChartForDeb.js";
import { MetricCard } from "./componets/metricCard.js";
import { StatsTable } from "./componets/statsTableForDebet.js";
import { appState } from "./state/appState.js";

export class DashboardRenderer {
    constructor(config) {
        this.config = config;
        this.components = [];
        this.statsSubscriptions = [];
    }

    build(container) {
        if (!container) return;
        container.innerHTML = '';

        // Очищаем старые подписки (если пересобираем)
        this.statsSubscriptions.forEach(unsub => unsub());
        this.statsSubscriptions = [];
        this.components = [];

        this.config.forEach(block => {
            const blockDiv = document.createElement('div');
            blockDiv.className = 'dashboard-block';

            const titleElem = document.createElement('h2');
            titleElem.textContent = block.title;
            blockDiv.appendChild(titleElem);

            if (block.type === 'charts') {
                const chartsContainer = document.createElement('div');
                chartsContainer.className = 'charts-grid';

                block.charts.forEach(chartConfig => {
                    const wrapperDiv = document.createElement('div');
                    wrapperDiv.className = 'chart-wrapper';

                    // Контейнер для столбчатого графика
                    const barContainer = document.createElement('div');
                    barContainer.style.width = '100%';
                    barContainer.style.height = '600px';
                    wrapperDiv.appendChild(barContainer);

                    const barChart = new ChartComponent(barContainer, chartConfig.metric, chartConfig.title);
                    this.components.push(barChart);

                    // Контейнер для круговой диаграммы с фоном
                    const pieWrapper = document.createElement('div');
                    pieWrapper.className = 'pie-wrapper';

                    const pieContainer = document.createElement('div');
                    pieContainer.style.width = '100%';
                    pieContainer.style.height = '350px';
                    pieWrapper.appendChild(pieContainer);

                    wrapperDiv.appendChild(pieWrapper);

                    const pieMetricKey = chartConfig.metric + 'Pie';
                    const pieTitle = `Распределение ${chartConfig.title.toLowerCase()}`;
                    const pieChart = new PieChartComponent(pieContainer, pieMetricKey, pieTitle);
                    this.components.push(pieChart);

                    chartsContainer.appendChild(wrapperDiv);
                });

                blockDiv.appendChild(chartsContainer);
            }
            else if (block.type === 'stats') {
                // Карточки статистики
                const statsContainer = document.createElement('div');
                statsContainer.className = 'stats-grid';
                const cards = [];

                block.metrics.forEach(metric => {
                    const cardContainer = document.createElement('div');
                    cardContainer.className = 'metric-card-wrapper';
                    statsContainer.appendChild(cardContainer);
                    const card = new MetricCard(cardContainer, metric.label, metric.key, metric.format);
                    cards.push(card);
                    this.components.push(card);
                });
                blockDiv.appendChild(statsContainer);

                // Контейнер для таблицы и графика (если оба есть)
                let rowContainer = null;
                let tableContainer = null;
                let chartContainer = null;

                // Создаём таблицу, если нужно
                if (block.table) {
                    tableContainer = document.createElement('div');
                    tableContainer.className = 'stats-table-container';
                    const tableComponent = new StatsTable(tableContainer);
                    this.components.push(tableComponent);
                    const unsubscribe = appState.subscribe((state) => {
                        if (state.tableData) {
                            tableComponent.render(state.tableData);
                        }
                    });
                    this.statsSubscriptions.push(unsubscribe);
                }

                // Создаём график, если нужно
                if (block.chart) {
                    chartContainer = document.createElement('div');
                    chartContainer.className = 'chart-wrapper-full';
                    chartContainer.style.width = '100%';
                    chartContainer.style.height = '470px';
                    const chart = new ChartComponent(chartContainer, block.chart.metric, block.chart.title);
                    this.components.push(chart);
                }

                // Если есть и таблица, и график — помещаем их в общий ряд
                if (tableContainer && chartContainer) {
                    rowContainer = document.createElement('div');
                    rowContainer.className = 'stats-row';
                    tableContainer.classList.add('stats-col');
                    chartContainer.classList.add('stats-col');
                    rowContainer.appendChild(tableContainer);
                    rowContainer.appendChild(chartContainer);
                    blockDiv.appendChild(rowContainer);
                } else if (tableContainer) {
                    blockDiv.appendChild(tableContainer);
                } else if (chartContainer) {
                    blockDiv.appendChild(chartContainer);
                }

                // Подписка на обновление карточек
                const unsubscribeStats = appState.subscribe((state) => {
                    if (state.stats) {
                        cards.forEach(card => card.update(state.stats));
                    }
                });
                this.statsSubscriptions.push(unsubscribeStats);
            }

            container.appendChild(blockDiv);
        });
    }
}