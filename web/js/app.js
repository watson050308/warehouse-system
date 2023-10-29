// 查看庫存按鈕點擊事件
document.getElementById("view-inventory").addEventListener("click", function() {
    // 在這裡處理查看庫存的相關操作，並將結果顯示在.results區域
    displayResults("查看庫存操作結果");
});

// 管理訂單按鈕點擊事件
document.getElementById("manage-orders").addEventListener("click", function() {
    // 在這裡處理管理訂單的相關操作，並將結果顯示在.results區域
    displayResults("管理訂單操作結果");
});

// 處理發貨按鈕點擊事件
document.getElementById("process-shipments").addEventListener("click", function() {
    // 在這裡處理處理發貨的相關操作，並將結果顯示在.results區域
    displayResults("處理發貨操作結果");
});

// 用於顯示結果的輔助函數
function displayResults(result) {
    document.querySelector(".results").innerHTML = result;
}