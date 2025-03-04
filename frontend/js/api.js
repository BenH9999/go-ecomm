document.addEventListener("DOMContentLoaded", () => {
    const loginForm = document.querySelector(".login-form");
    if(loginForm){
        loginForm.addEventListener("submit", async (e) => {
            e.preventDefault();
            const email=document.getElementById("email").value;
            const password= document.getElementById("password").value;
            const response = await fetch("/api/login", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({email, password})
            });
            const result = await response.json();
            console.log(result.username);
            if(response.ok){
                alert(result.message + " (" + result.username + ")");
                window.location.href = "/";
            }
            else{
                alert(result.message || "Login failed");
            }
        });
    }
});

document.addEventListener("DOMContentLoaded", () => {
    const registerForm=document.querySelector(".register-form");
    if(registerForm){
        registerForm.addEventListener("submit", async (e) => {
            e.preventDefault();
            const username=document.getElementById("username").value;
            const email=document.getElementById("email").value;
            const password=document.getElementById("password").value;
            const response = await fetch("/api/register", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({username, email, password})
            });
            const result = await response.json();
            if(response.ok){
                alert(result.message + "registered: " + result.username);
                window.location.href = "login.html";
            }
            else{
                alert(result.message || "Register failed");
            }
        });
    }
});
