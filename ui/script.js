const sendBtn = document.getElementById("send-btn");
const userInput = document.getElementById("user-input");
const chatBox = document.getElementById("chat-box");

sendBtn.addEventListener("click", async () => {
    const message = userInput.value;
    if (!message) return;

    // Display user message
    chatBox.innerHTML += `<div><strong>You:</strong> ${message}</div>`;
    userInput.value = "";

    // Send message to the backend
    const response = await fetch("http://localhost:8080/chat", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ message }),
    });

    const data = await response.json();

    // Display chatbot response
    chatBox.innerHTML += `<div><strong>Bot:</strong> ${data.message}</div>`;
    chatBox.scrollTop = chatBox.scrollHeight;
});
