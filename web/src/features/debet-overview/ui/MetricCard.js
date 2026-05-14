export class MetricCard {
    constructor(container, title, valueKey, format = 'number') {
        this.container = container;
        this.title = title;
        this.valueKey = valueKey;
        this.format = format;
        this.valueElement = null;
        this.render(null);
    }

    update(stats) {
        if (!stats) return;
        const value = stats[this.valueKey];
        let formattedValue = '—';

        if (value !== undefined && value !== null) {
            if (this.format === 'number') {
                formattedValue = value.toLocaleString();
            } else if (this.format === 'currency') {
                const inBillions = value / 1_000_000_000;
                const rounded = Math.round(inBillions * 10) / 10;
                formattedValue = rounded.toLocaleString('ru-RU').replace('.', ',') + ' млрд ₽';
            } else {
                formattedValue = value;
            }
        }

        if (this.valueElement) {
            this.valueElement.textContent = formattedValue;
        }
    }

    render(stats) {
        this.container.innerHTML = '';
        const cardDiv = document.createElement('div');
        cardDiv.className = 'metric-card';

        const titleDiv = document.createElement('div');
        titleDiv.className = 'card-title';
        titleDiv.textContent = this.title;
        cardDiv.appendChild(titleDiv);

        const valueDiv = document.createElement('div');
        valueDiv.className = 'card-value';
        valueDiv.textContent = 'Загрузка...';
        cardDiv.appendChild(valueDiv);

        this.container.appendChild(cardDiv);
        this.valueElement = valueDiv;

        if (stats) this.update(stats);
    }
}
