document.getElementById("fetchData").addEventListener("click", async () => {
    const res = await fetch("/api/test");
    const text = await res.text();
    document.getElementById("response").innerText = text;
});
