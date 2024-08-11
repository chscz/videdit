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

function addRow() {
   var container = document.getElementById("edit_container");
   var rowCount = container.children.length;

   var row = document.createElement("div");
   row.className = "row";

   row.innerHTML = `
       <select class="video_select">
           ${Array.from(document.querySelectorAll("#video_select option"))
              .map(
                 (option) =>
                    `<option value="${option.value}">${option.textContent}</option>`
              )
              .join("")}
       </select>
       <input type="number" placeholder="숫자 입력" />
       <input type="number" placeholder="숫자 입력" />
       <button onclick="removeRow(this)">삭제</button>
   `;

   container.appendChild(row);
}

function removeRow(button) {
   var row = button.parentElement;
   row.remove();
}

function createVideo() {
   var rows = document.querySelectorAll("#edit_container .row");

   var extensionSelect = document.getElementById("extension_select");
   var ext = extensionSelect.value;

   var data = Array.from(rows)
      .map((row, index) => {
         var select = row.querySelector("select.video_select");
         var inputs = row.querySelectorAll("input[type='number']");

         if (!select || inputs.length < 2) {
            console.error("필수 요소가 누락되었습니다.");
            return null;
         }
         console.log(index);
         return {
            order: index || 0,
            videoId: select.value || "",
            videoFileName:
               select.options[select.selectedIndex].textContent || "",
            trimStart: parseFloat(inputs[0].value) || -1,
            trimEnd: parseFloat(inputs[1].value) || -1,
         };
      })
      .filter((item) => item !== null);

   var requestData = {
      ext: ext,
      videos: data,
   };

   fetch("/create_video", {
      method: "POST",
      headers: {
         "Content-Type": "application/json",
      },
      body: JSON.stringify(requestData),
   })
      .then((response) => response.json())
      .then((result) => {
         console.log("서버 응답:", result);
      })
      .catch((error) => {
         console.error("오류:", error);
      });
}
