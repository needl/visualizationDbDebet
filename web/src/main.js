// src/main.js
import { dashboardConfig } from './dashboardConfig.js'
import { loadData } from './state/actionForDebet.js'
import { loadStats } from './state/actionForResponse.js'
import { DashboardRenderer } from './dashboardRender.js'

const renderer = new DashboardRenderer(dashboardConfig);
const app = document.getElementById('app');
if (app) {
    renderer.build(app);
}

// Загружаем данные
loadStats();
loadData();