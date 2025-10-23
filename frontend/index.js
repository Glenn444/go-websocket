
window.onload = function(){
    
    const messages = document.getElementById("messages")
    const myform = document.getElementById("form")

    myform.addEventListener("submit",function(e){
        e.preventDefault()
        console.log("submitted form");
        const formData = new FormData(myform)
        

        const message = formData.get('message')
        const chatroom = formData.get('chatroom')


        console.log("message: ",message);
        console.log("chatroom: ",chatroom);
        
    })
    console.log("window loaded");
    
}


