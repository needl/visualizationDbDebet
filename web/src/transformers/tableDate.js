// src/transformers/tableData.js
export function prepareTableData(aggregated) {
    const rows = [];
    for (const [name, data] of Object.entries(aggregated)) {
        const contractAmount = data.contractAmount || 0;
        const debetTotal = data.debetTotal || 0;
        const coefficient = contractAmount > 0 ? debetTotal / contractAmount : 0;
        rows.push({
            name,
            contractAmount,
            debetTotal,
            coefficient
        });
    }
    rows.sort((a, b) => b.coefficient - a.coefficient);
    return rows;
}