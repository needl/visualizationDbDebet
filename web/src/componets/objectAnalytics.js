// src/components/ObjectAnalytics.js
import { appState } from '../state/appState.js';
import { ObjectMetricCard } from './objectMetricCard.js';

export class ObjectAnalytics {
    constructor(container) {
        this.container = container;
        this.cards = [];
        this.chart = null;
        this.unsubscribe = null;
        this.mount();
    }

    mount() {
        this.render();
        this.subscribe();
    }

    render() {
        this.container.innerHTML = `
            <div class="object-analytics">
                <div class="object-metrics-grid"></div>
                <div class="analytics-chart" style="width:100%; height:400px;"></div>
            </div>
        `;
    }

    subscribe() {
        this.unsubscribe = appState.subscribe(state => {
            if (state.objectData && state.objectData.metrics) {
                this.updateCards(state.objectData.metrics);
                this.updateChart(state.objectData.chartData);
            } else {
                this.clearData();
            }
        });
    }

    updateCards(metrics) {
        const container = this.container.querySelector('.object-metrics-grid');
        if (!container) return;
        container.innerHTML = '';
        this.cards = [];

        // 10 карточек в порядке, как требовалось (без Name)
        const metricDefs = [
            { label: 'Подрядчик', key: 'contractorName', format: 'string' },
            { label: 'Дата начала работ', key: 'workStartDate', format: 'date' },
            { label: 'Дата окончания работ', key: 'workEndDate', format: 'date' },
            { label: 'Строительная готовность', key: 'buildReadyPercent', format: 'boolean' },
            { label: 'Разрешение на ввод', key: 'permissionToEnter', format: 'boolean' },
            { label: 'Заключение МКЭ', key: 'conclusionMke', format: 'boolean' },
            { label: 'Твёрдая договорная цена', key: 'hardContractPrice', format: 'money' },
            { label: 'Сумма договора', key: 'contractAmount', format: 'money' },
            { label: 'Перечислено', key: 'paidAmount', format: 'money' },
            { label: 'Принято', key: 'acceptedAmount', format: 'money' },
        ];

        metricDefs.forEach(def => {
            const wrapper = document.createElement('div');
            wrapper.className = 'metric-card-wrapper';
            const cardDiv = document.createElement('div');
            wrapper.appendChild(cardDiv);
            container.appendChild(wrapper);

            const card = new ObjectMetricCard(cardDiv, def.label, def.key, def.format);
            card.update(metrics);
            this.cards.push(card);
        });
    }

    updateChart(chartData) {
        const chartContainer = this.container.querySelector('.analytics-chart');
        if (!chartContainer) return;

        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }

        if (!chartData || !chartData.series || chartData.series.length === 0) {
            return;
        }

        this.chart = echarts.init(chartContainer);

        const option = {
            title: { text: 'График задолженности по годам', left: 'center' },
            tooltip: { trigger: 'axis' },
            legend: { data: chartData.series.map(s => s.name), bottom: 0 },
            xAxis: { type: 'category', data: chartData.categories },
            yAxis: {
                type: 'value',
                axisLabel: {
                    formatter: (val) => (val / 1_000_000).toFixed(1) + ' млн'
                }
            },
            series: chartData.series
        };

        this.chart.setOption(option, true);
        this.chart.resize();
    }

    clearData() {
        const cardsContainer = this.container.querySelector('.object-metrics-grid');
        if (cardsContainer) cardsContainer.innerHTML = '';
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
    }

    destroy() {
        if (this.unsubscribe) this.unsubscribe();
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
    }
}