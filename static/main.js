const input = document.querySelector('#textarea')
const messages = document.querySelector('#messages')
const username = document.querySelector('#username')
const enter = document.querySelector('#enter-name')
const send = document.querySelector('#send')

let socket;
// let socket = new WebSocket("ws://localhost:8080/chat?username=wael");
// socket.onmessage = function(event) {
//   console.log(`[message] Data received from server: ${event.data}`);
// };

enter.addEventListener("click", () => {
  username.value
  socket = new WebSocket(`ws://localhost:8080/chat?username=${username}`);
  socket.onclose = () => {
    socket.close()
  };
  username.disabled = true
  enter.disabled = true

  socket.onmessage = function (msg) {
    insertMessage(JSON.parse(msg.data))
  };

  send.onclick = () => {
    const message = {
      "command": 2,
      "channel": "general",
      "content": input.value,
      "username": username.value,
    }

    socket.send(JSON.stringify(message));
    input.value = "";
  };

})

/**
 * Insert a message into the UI
 * @param {Message that will be displayed in the UI} messageObj
 */
function insertMessage(messageObj) {
  // Create a div object which will hold the message
  const message = document.createElement('div')
  messageObj.content.split("--")

  // Set the attribute of the message div
  message.setAttribute('class', 'chat-message')
  message.textContent = `${messageObj.content.split("--")[0]}: ${messageObj.content.split("--")[1]}`

  // Append the message to our chat div
  messages.appendChild(message)

  // Insert the message as the first message of our chat
  messages.insertBefore(message, messages.firstChild)
}