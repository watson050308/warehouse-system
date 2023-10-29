// 註冊表單提交事件
document.getElementById("register-form").addEventListener("submit", function(e) {
    e.preventDefault();
    const newEmail = document.getElementById("new-email").value;
    const newPassword = document.getElementById("new-password").value;

    // 在這裡處理註冊操作，可以向後端發送註冊請求
    console.log("使用者嘗試註冊：");
    console.log("新電子郵件：" + newEmail);
    console.log("新密碼：" + newPassword);
});