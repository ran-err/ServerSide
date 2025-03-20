let captchaID;

document.addEventListener("DOMContentLoaded", function () {
    document.getElementById("verifyCaptcha").addEventListener("click", verifyCaptcha);
    loadCaptcha();
});

function loadCaptcha() {
    fetch("captcha/new")
        .then(response => response.json())
        .then(data => {
            captchaID = data.captcha_id;
            document.getElementById("captchaImage").src = data.captcha_image;
        });
}

function verifyCaptcha() {
    const userInput = document.getElementById("captchaInput").value;

    fetch("captcha/verify", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({captcha_id: captchaID, captcha_solution: userInput})
    })
        .then(response => response.json())
        .then(data => {
            document.getElementById("captchaStatus").innerText = data.message;
            if (data.success) loadCaptcha();
        });
}
