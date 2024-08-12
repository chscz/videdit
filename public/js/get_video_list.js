function getVideoList() {
   fetch("/get_video_list")
      .then((response) => response.json())
      .then((data) => {
         var videoListContainer = document.getElementById("video_list");
         videoListContainer.innerHTML = "";

         // 업로드된 동영상 목록
         var uploadVideosTitle = document.createElement("h3");
         uploadVideosTitle.textContent = "업로드된 동영상:";
         videoListContainer.appendChild(uploadVideosTitle);
         console.log(data.uploadVideos);
         var uploadVideosList = document.createElement("ul");
         data.uploadVideos.forEach((video) => {
            var listItem = document.createElement("li");
            listItem.textContent = `ID: ${video.id},
             생성일: ${new Date(video.created_at).toLocaleString()},
              파일명: ${video.file_name},
               경로: ${video.file_path}`;
            uploadVideosList.appendChild(listItem);
         });
         videoListContainer.appendChild(uploadVideosList);
         console.log(data.uploadVideos);
         // 생성된 동영상 목록
         var createVideosTitle = document.createElement("h3");
         createVideosTitle.textContent = "생성된 동영상:";
         videoListContainer.appendChild(createVideosTitle);

         var createVideosList = document.createElement("ul");
         data.reqVideos.forEach((video) => {
            var listItem = document.createElement("li");
            listItem.textContent = `ID: ${video.id},
             생성일: ${new Date(video.created_at).toLocaleString()},
              요청: ${video.request},
              경로: ${video.file_path}`;
            createVideosList.appendChild(listItem);
         });
         videoListContainer.appendChild(createVideosList);
      })
      .catch((error) => {
         console.error("동영상 목록 조회 실패:", error);
      });
}
