window.onload = function () {
  const messages = document.getElementById("messages");
  const myform = document.getElementById("form");
  let socket;

  if (window["WebSocket"]) {
    socket = new WebSocket("ws://" + document.location.host + "/ws");
    console.log(socket);

    console.log("supports websockets");
  } else {
    this.alert("this browser does not support websockets");
  }
  myform.addEventListener("submit", function (e) {
    e.preventDefault();
    console.log("submitted form");
    const formData = new FormData(myform);

    const message = formData.get("message");
    const chatroom = formData.get("chatroom");

    console.log("socket: ",socket);
    
    if (socket.readState == 1) {
      console.log("ready state");
      socket.send(message);
    }

    console.log("message: ", message);
    console.log("chatroom: ", chatroom);
  });
  console.log("window loaded");
};
