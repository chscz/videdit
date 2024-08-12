function uploadFile(event) {
   event.preventDefault();

   var formData = new FormData();
   var fileField = document.querySelector("#upload_file");

   if (!fileField.files.length) {
      return;
   }

   formData.append("upload_file", fileField.files[0]);

   fetch("/upload", {
      method: "POST",
      body: formData,
   })
      .then((response) => response.json())
      .then((data) => {
         if (data) {
            var list = document.getElementById("upload_list");
            var newItem = document.createElement("li");
            newItem.textContent = `ID: ${data.id}, 파일명: ${data.file_name}`;
            newItem.dataset.id = data.id;
            list.appendChild(newItem);

            var selectElement = document.getElementById("video_select");
            var option = document.createElement("option");
            option.value = data.id;
            option.textContent = data.file_name;
            selectElement.appendChild(option);

            fileField.value = "";
         }
      })
      .catch((error) => {
         console.error("Error:", error);
      });
}
