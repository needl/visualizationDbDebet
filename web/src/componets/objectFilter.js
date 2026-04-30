// src/components/ObjectFilter.js
import { appState } from '../state/appState.js';
import { loadObjectList } from '../state/actions/actionForObjects.js';
import { loadObjectData } from '../state/actions/actionForObjects.js';

export class ObjectFilter {
    constructor(container) {
        this.container = container;
        this.customSelect = null;
        this.selectedValueDiv = null;
        this.optionsList = null;
        this.unsubscribe = null;
        this._currentCustomer = null;
        this.init();
    }

    init() {
        this.render(); // начальная структура
        this.subscribe();
    }

    subscribe() {
        this.unsubscribe = appState.subscribe(state => {
            const { selectedCustomer, objectList } = state;
            if (!selectedCustomer) {
                this.clear();
            } else if (this._currentCustomer !== selectedCustomer) {
                this._currentCustomer = selectedCustomer;
                this.loadAndRenderOptions(selectedCustomer);
            }
        });
    }

    async loadAndRenderOptions(orgName) {
        this.showLoading();
        await loadObjectList(orgName);
        this.renderOptions();
    }

    render() {
        this.container.innerHTML = ''; // скрыто, пока не выбран заказчик
        this.customSelect = null;
        this.optionsList = null;
        this.selectedValueDiv = null;
    }

    showLoading() {
        this.container.innerHTML = '<div class="object-filter-loading">Загрузка объектов...</div>';
    }

    renderOptions() {
        const state = appState._getState();
        const objectList = state.objectList || [];
        if (!objectList.length) {
            this.container.innerHTML = '<div class="object-filter-empty">Нет доступных объектов</div>';
            return;
        }

        // Создаём кастомный селект
        const filterWrapper = document.createElement('div');
        filterWrapper.className = 'object-filter';

        const label = document.createElement('label');
        label.textContent = 'Выберите объект:';
        label.htmlFor = 'object-custom-select';
        filterWrapper.appendChild(label);

        const customSelect = document.createElement('div');
        customSelect.className = 'custom-select';

        // Отображаемое значение (после выбора – только первое слово)
        const selectedValue = document.createElement('div');
        selectedValue.className = 'selected-value';
        selectedValue.textContent = '-- Объект --';
        selectedValue.addEventListener('click', () => {
            // Переключаем видимость списка
            const options = customSelect.querySelector('.options');
            if (options) options.classList.toggle('open');
        });

        // Выпадающий список
        const optionsList = document.createElement('ul');
        optionsList.className = 'options';
        objectList.forEach(name => {
            const li = document.createElement('li');
            li.className = 'option-item';
            li.textContent = name;
            li.addEventListener('click', () => {
                // Обновляем отображаемое значение (первое слово)
                const firstWord = name.split(' ')[0] || name;
                selectedValue.textContent = firstWord;
                // Закрываем список
                optionsList.classList.remove('open');

                // Вызываем загрузку данных
                const orgName = appState._getState().selectedCustomer;
                appState.setSelectedObject(name); // сохраняем полное имя в стейт
                loadObjectData(orgName, name);
            });
            optionsList.appendChild(li);
        });

        // Закрытие списка при клике вне
        document.addEventListener('click', (e) => {
            if (!customSelect.contains(e.target)) {
                optionsList.classList.remove('open');
            }
        });

        customSelect.appendChild(selectedValue);
        customSelect.appendChild(optionsList);
        filterWrapper.appendChild(customSelect);
        this.container.innerHTML = '';
        this.container.appendChild(filterWrapper);

        this.customSelect = customSelect;
        this.selectedValueDiv = selectedValue;
        this.optionsList = optionsList;
    }

    clear() {
        this._currentCustomer = null;
        this.container.innerHTML = '';
        this.customSelect = null;
        this.optionsList = null;
        this.selectedValueDiv = null;
    }
}