let rowCounter = 2;

function addRow() {
   var container = document.getElementById("edit_container");
   var rowCount = container.children.length;

   var row = document.createElement("div");
   row.className = "row";

   row.innerHTML = `
      <span class="row-number">${rowCounter}</span>
      <span class="row-symbol">></span>
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

   rowCounter++;
}

function removeRow(button) {
   var row = button.parentElement;
   row.remove();

   reorderRows();
}

function reorderRows() {
   var container = document.getElementById("edit_container");
   var rows = container.querySelectorAll(".row");

   rows.forEach((row, index) => {
      row.querySelector(".row-number").textContent = index + 1;
   });

   rowCounter = rows.length + 1;
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

         return {
            order: index || 0,
            video_id: select.value || "",
            video_file_name:
               select.options[select.selectedIndex].textContent || "",
            trim_start: parseFloat(inputs[0].value) || -1,
            trim_end: parseFloat(inputs[1].value) || -1,
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
         if (result && result.file_name && result.url) {
            addDownloadLink(result.file_name, result.url);
         }
      })
      .catch((error) => {
         console.error("오류:", error);
      });
}

function addDownloadLink(filename, url) {
   var downloadLinksContainer = document.getElementById("download_links");

   var linkRow = document.createElement("div");
   linkRow.className = "link-row";

   var downloadText = document.createElement("span");
   downloadText.textContent = "URL: ";

   var downloadLink = document.createElement("a");
   downloadLink.href = url;
   downloadLink.textContent = url;
   downloadLink.target = "_blank";

   var downloadButton = document.createElement("button");
   downloadButton.textContent = "다운로드";
   downloadButton.onclick = function () {
      window.location.href = url;
   };

   var copyButton = document.createElement("button");
   copyButton.textContent = "URL 링크 복사";
   copyButton.onclick = function () {
      navigator.clipboard
         .writeText(url)
         .then(() => {
            alert("URL이 클립보드에 복사되었습니다.");
         })
         .catch((err) => {
            console.error("클립보드 복사 실패:", err);
         });
   };

   linkRow.appendChild(downloadText);
   linkRow.appendChild(downloadLink);
   linkRow.appendChild(downloadButton);
   linkRow.appendChild(copyButton); // 복사 버튼 추가
   downloadLinksContainer.appendChild(linkRow);
}
