// src/components/objectMetricCard.js
export class ObjectMetricCard {
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
        let value = stats[this.valueKey];
        let formattedValue = '—';

        if (value !== undefined && value !== null) {
            switch (this.format) {
                case 'number':
                    formattedValue = value.toLocaleString();
                    break;
                case 'money':
                    const inMillions = value / 1_000_000;
                    const rounded = Math.round(inMillions * 10) / 10;
                    formattedValue = rounded.toLocaleString('ru-RU').replace('.', ',') + ' млн ₽';
                    break;
                case 'percent':
                    formattedValue = value.toFixed(2).replace('.', ',') + '%';
                    break;
                case 'date':
                    formattedValue = value;
                    break;
                case 'boolean':
                    formattedValue = value ? 'Да' : 'Нет';
                    break;
                case 'string':
                default:
                    formattedValue = String(value);
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
        valueDiv.textContent = '—';
        cardDiv.appendChild(valueDiv);
        this.container.appendChild(cardDiv);
        this.valueElement = valueDiv;

        if (stats) this.update(stats);
    }
}