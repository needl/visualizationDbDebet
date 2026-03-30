// src/main.js
import { dashboardConfig } from './dashboardConfig.js'
import { loadData } from './state/actions/actionForDebet.js'
import { loadStats } from './state/actions/actionForResponse.js'
import {loadCustomerData} from "./state/actions/actionForCustomers.js";
import { DashboardRenderer } from './dashboardRender.js'

const renderer = new DashboardRenderer(dashboardConfig);
const app = document.getElementById('app');
if (app) {
    renderer.build(app);
}

// Загружаем данные
loadStats();
loadData();
loadCustomerData();
