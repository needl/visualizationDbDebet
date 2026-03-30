export class BlockFactorsChart {
    constructor(container, title = 'Блок-факторы по количеству подрядчиков') {
        this.container = container;
        this.title = title;
        this.chart = null;
    }

    render(blockFactors) {
        if (!blockFactors) {
            this.clear();
            this.container.innerHTML = '<div class="empty-message">Нет данных по блок-факторам</div>';
            return;
        }

        const data = Object.entries(blockFactors)
            .map(([name, value]) => ({ name, value }))
            .filter(item => item.value > 0);

        if (data.length === 0) {
            this.clear();
            this.container.innerHTML = '<div class="empty-message">Нет данных по блок-факторам</div>';
            return;
        }

        if (!this.chart) {
            this.chart = echarts.init(this.container);
        }

        const names = data.map(item => this.formatName(item.name));
        const values = data.map(item => item.value);

        const option = {
            title: { text: this.title, left: 'center' },
            tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
            grid: { containLabel: true, left: '20%' },
            xAxis: { type: 'value', name: 'Количество' },
            yAxis: { type: 'category', data: names, axisLabel: { fontSize: 10 } },
            series: [{
                name: 'Количество',
                type: 'bar',
                data: values,
                itemStyle: { color: '#d11b1b' },
                label: { show: true, position: 'right' }
            }]
        };

        this.chart.setOption(option, true);
        this.chart.resize(); // адаптация к размерам контейнера
    }

    clear() {
        if (this.chart) {
            this.chart.dispose();
            this.chart = null;
        }
        this.container.innerHTML = '';
        // высота не сбрасывается – она задана в CSS
    }

    formatName(key) {
        const map = {
            bankrot_count: 'Банкротство',
            likvidatsiya_count: 'Ликвидация',
            nedostovernost_count: 'Недостоверность данных',
            isklyuchenie_count: 'Исключение из ЕГРЮЛ',
            inostrannye_count: 'Реестр иностранных агентов',
            eks_ter_count: 'Экстремизм Терроризм',
            nedobrosovestn_count: 'Реестр недобросовестных поставщиков',
            admin_otvet_count: 'Административные правонарушения',
            nam_bankrot_count: 'Намерение о банкротство',
            blokirovka_count: 'Блокировка счётов',
            chisl_count: 'Численность меньше 1'
        };
        return map[key] || key;
    }
}