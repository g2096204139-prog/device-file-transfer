const tokenInput = document.getElementById("token");
const saveTokenBtn = document.getElementById("saveTokenBtn");
const fileInput = document.getElementById("fileInput");
const uploadBtn = document.getElementById("uploadBtn");
const refreshBtn = document.getElementById("refreshBtn");
const uploadProgress = document.getElementById("uploadProgress");
const progressText = document.getElementById("progressText");
const uploadStatus = document.getElementById("uploadStatus");
const fileList = document.getElementById("fileList");
const serverInfo = document.getElementById("serverInfo");

function getToken() {
  return localStorage.getItem("dft_token") || "";
}

function setStatus(message) {
  uploadStatus.textContent = typeof message === "string"
    ? message
    : JSON.stringify(message, null, 2);
}

function resetUploadInput() {
  fileInput.value = "";
  uploadProgress.value = 0;
  progressText.textContent = "0%";
}

function formatBytes(bytes) {
  if (bytes === 0) return "0 B";
  const units = ["B", "KB", "MB", "GB", "TB"];
  const index = Math.floor(Math.log(bytes) / Math.log(1024));
  const value = bytes / Math.pow(1024, index);
  return `${value.toFixed(value >= 10 ? 1 : 2)} ${units[index]}`;
}

async function fetchJSON(url, options = {}) {
  const headers = options.headers || {};
  headers["X-Access-Token"] = getToken();

  const response = await fetch(url, {
    ...options,
    headers,
  });

  const data = await response.json().catch(() => ({}));

  if (!response.ok || data.success === false) {
    throw new Error(data.message || `Request failed: ${response.status}`);
  }

  return data;
}

async function loadServerInfo() {
  try {
    const result = await fetchJSON("/api/server-info", { headers: {} });
    const data = result.data;
    const lan = (data.lan_ips || []).map(ip => `http://${ip}:${data.port}`).join(" ｜ ");
    serverInfo.textContent = `Version ${data.version} ｜ Max upload: ${data.max_upload_size_mb} MB ｜ LAN URLs: ${lan || "No LAN IP detected"}`;
  } catch (error) {
    serverInfo.textContent = "Unable to load server info.";
  }
}

async function loadFiles() {
  fileList.innerHTML = "Loading...";

  try {
    const result = await fetchJSON("/api/files");
    const files = result.data || [];

    if (files.length === 0) {
      fileList.innerHTML = "<p>No files uploaded yet.</p>";
      return;
    }

    fileList.innerHTML = "";

    for (const file of files) {
      const item = document.createElement("div");
      item.className = "file-item";

      const info = document.createElement("div");
      info.innerHTML = `
        <strong>${escapeHTML(file.filename)}</strong>
        <div class="file-meta">${formatBytes(file.size)} ｜ ${new Date(file.uploaded_at).toLocaleString()}</div>
      `;

      const download = document.createElement("button");
      download.type = "button";
      download.textContent = "Download";
      download.onclick = () => downloadFile(file.filename);

      const remove = document.createElement("button");
      remove.type = "button";
      remove.textContent = "Delete";
      remove.className = "danger";
      remove.onclick = () => deleteFile(file.filename);

      item.appendChild(info);
      item.appendChild(download);
      item.appendChild(remove);
      fileList.appendChild(item);
    }
  } catch (error) {
    fileList.innerHTML = `<p>${escapeHTML(error.message)}</p>`;
  }
}

function downloadFile(filename) {
  const url = `/api/download/${encodeURIComponent(filename)}`;

  fetch(url, {
    headers: { "X-Access-Token": getToken() }
  })
    .then(response => {
      if (!response.ok) throw new Error("Download failed");
      return response.blob();
    })
    .then(blob => {
      const objectUrl = URL.createObjectURL(blob);
      const link = document.createElement("a");
      link.href = objectUrl;
      link.download = filename;
      link.target = "_blank";
      link.rel = "noopener";

      document.body.appendChild(link);
      link.click();
      link.remove();

      setTimeout(() => URL.revokeObjectURL(objectUrl), 60000);
      setStatus(`Download started: ${filename}`);
    })
    .catch(error => alert(error.message));
}

async function deleteFile(filename) {
  const confirmed = confirm(`Delete ${filename}?`);
  if (!confirmed) return;

  try {
    await fetchJSON(`/api/files?filename=${encodeURIComponent(filename)}`, {
      method: "DELETE",
    });
    setStatus(`Deleted: ${filename}`);
    resetUploadInput();
    await loadFiles();
  } catch (error) {
    alert(error.message);
  }
}

function uploadFiles() {
  const files = fileInput.files;
  if (!files.length) {
    setStatus("Please choose at least one file.");
    return;
  }

  const formData = new FormData();
  for (const file of files) {
    formData.append("files", file);
  }

  const xhr = new XMLHttpRequest();
  xhr.open("POST", "/api/upload");
  xhr.setRequestHeader("X-Access-Token", getToken());

  xhr.upload.onprogress = event => {
    if (!event.lengthComputable) return;
    const percent = Math.round((event.loaded / event.total) * 100);
    uploadProgress.value = percent;
    progressText.textContent = `${percent}%`;
  };

  xhr.onload = async () => {
    try {
      const data = JSON.parse(xhr.responseText || "{}");
      setStatus(data);
      resetUploadInput();
      await loadFiles();
    } catch {
      setStatus(xhr.responseText);
    }
  };

  xhr.onerror = () => {
    setStatus("Upload failed.");
  };

  uploadProgress.value = 0;
  progressText.textContent = "0%";
  setStatus("Uploading...");
  xhr.send(formData);
}

function escapeHTML(value) {
  return String(value).replace(/[&<>"']/g, char => ({
    "&": "&amp;",
    "<": "&lt;",
    ">": "&gt;",
    '"': "&quot;",
    "'": "&#39;",
  }[char]));
}

saveTokenBtn.onclick = () => {
  localStorage.setItem("dft_token", tokenInput.value);
  setStatus("Token saved.");
  loadFiles();
};

uploadBtn.onclick = uploadFiles;
refreshBtn.onclick = loadFiles;

tokenInput.value = getToken();
loadServerInfo();
loadFiles();
