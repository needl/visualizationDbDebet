import { dashboardConfig } from './dashboardConfig.js';
import { loadData } from './state/actions/actionForDebet.js';
import { loadStats } from './state/actions/actionForResponse.js';
import { DashboardRenderer } from './dashboardRender.js';

const renderer = new DashboardRenderer(dashboardConfig);
const app = document.getElementById('app');

if (app) {
    renderer.build(app);
}

loadStats();
loadData();

