export class StatsTable {
    constructor(container) {
        this.container = container;
    }

    render(data) {
        if (!data || !data.length) {
            this.container.innerHTML = '<div class="empty">Нет данных для таблицы</div>';
            return;
        }

        const table = document.createElement('table');
        table.className = 'stats-table';

        const thead = document.createElement('thead');
        thead.innerHTML = `
            <tr>
                <th>Заказчик</th>
                <th>Цена контрактов</th>
                <th>Дебиторская задолженность</th>
                <th>Соотношение дебиторской задолженности к сумме контрактов</th>
            </tr>
        `;
        table.appendChild(thead);

        const tbody = document.createElement('tbody');
        data.forEach(row => {
            const tr = document.createElement('tr');

            const tdName = document.createElement('td');
            tdName.textContent = row.name;
            tr.appendChild(tdName);

            const tdContract = document.createElement('td');
            tdContract.textContent = (row.contractAmount / 1_000_000_000).toLocaleString('ru-RU') + ' млрд ₽';
            tr.appendChild(tdContract);

            const tdDebet = document.createElement('td');
            tdDebet.textContent = (row.debetTotal / 1_000_000_000).toLocaleString('ru-RU') + ' млрд ₽';
            tr.appendChild(tdDebet);

            const tdCoeff = document.createElement('td');
            const percent = Math.min(row.coefficient, 100);
            tdCoeff.innerHTML = `
                <div class="progress-container">
                    <div class="progress-bar">
                        <div class="progress-fill" style="width: ${percent}%;"></div>
                    </div>
                    <span class="coeff-value">${row.coefficient.toFixed(1) + '%'}</span>
                </div>
            `;
            tr.appendChild(tdCoeff);

            tbody.appendChild(tr);
        });
        table.appendChild(tbody);

        this.container.innerHTML = '';
        this.container.appendChild(table);
    }
}
