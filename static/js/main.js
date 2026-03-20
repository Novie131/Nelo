// static/js/main.js

let stagedFiles = [];

function clearHub() {
    if (!confirm("確定要清空 Hub 裡的專案資料嗎？")) return;
    fetch('/api/context', { method: 'DELETE' })
    .then(response => {
        if (response.ok) window.location.reload();
        else alert("清空失敗");
    }).catch(err => alert("❌ 無法連線至伺服器"));
}

function stageMarkdown() {
    const content = document.getElementById('markdown-content').textContent;
    const btn = document.getElementById('stage-btn');
    const baseName = btn ? btn.getAttribute('data-project-name') : 'project';
    
    const date = new Date();
    const timeString = `${date.getHours()}-${date.getMinutes()}-${date.getSeconds()}`;
    const fileName = `${baseName}_${timeString}.md`;

    stagedFiles.push({
        id: Date.now().toString(),
        name: fileName,
        content: content
    });

    renderStagingArea();
    alert(`已將 ${fileName} 加入準備區！`);
}

function renderStagingArea() {
    const fileList = document.getElementById('file-list');
    const emptyMsg = document.getElementById('empty-staging-msg');

    // 如果沒有檔案，顯示空狀態提示
    if (stagedFiles.length === 0) {
        fileList.innerHTML = '';
        fileList.classList.add('hidden');
        emptyMsg.classList.remove('hidden');
        return;
    }

    // 如果有檔案，顯示列表
    emptyMsg.classList.add('hidden');
    fileList.classList.remove('hidden');
    fileList.innerHTML = '';

    stagedFiles.forEach(file => {
        const row = document.createElement('label');
        row.className = 'flex items-center gap-3 p-3 bg-gray-900 rounded-lg border border-gray-700 cursor-pointer hover:border-blue-500 transition-colors';
        const sizeKB = (new Blob([file.content]).size / 1024).toFixed(1);

        row.innerHTML = `
            <input type="checkbox" value="${file.id}" checked class="file-checkbox w-5 h-5 text-blue-600 rounded border-gray-600 bg-gray-800 cursor-pointer">
            <span class="text-gray-300 font-mono text-sm select-none flex-grow">${file.name}</span>
            <span class="text-xs text-gray-500 font-mono bg-gray-800 px-2 py-1 rounded">${sizeKB} KB</span>
        `;
        fileList.appendChild(row);
    });
}

function getSelectedFiles() {
    const checkboxes = document.querySelectorAll('.file-checkbox:checked');
    const selectedIds = Array.from(checkboxes).map(cb => cb.value);
    return stagedFiles.filter(f => selectedIds.includes(f.id));
}

function downloadSelected() {
    const selected = getSelectedFiles();
    if (selected.length === 0) return alert('⚠️ 請先勾選要下載的檔案！');

    selected.forEach(file => {
        const blob = new Blob([file.content], { type: 'text/markdown;charset=utf-8' });
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = file.name;
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
        URL.revokeObjectURL(url);
    });
}

function copySelected() {
    const selected = getSelectedFiles();
    if (selected.length === 0) return alert('⚠️ 請先勾選要複製的檔案！');

    const combinedContent = selected.map(f => f.content).join('\n\n========================================\n\n');
    executeCopy(combinedContent, `✅ 成功複製純文字內容！可以直接在 Gemini 對話框按 Ctrl+V 貼上了。`);
}

function deleteSelected() {
    const selectedIds = Array.from(document.querySelectorAll('.file-checkbox:checked')).map(cb => cb.value);
    if (selectedIds.length === 0) return alert('⚠️ 請先勾選要移除的檔案！');

    stagedFiles = stagedFiles.filter(f => !selectedIds.includes(f.id));
    renderStagingArea();
}

function executeCopy(text, successMsg) {
    if (navigator.clipboard && window.isSecureContext) {
        navigator.clipboard.writeText(text)
            .then(() => alert(successMsg))
            .catch(err => fallbackCopy(text, successMsg));
    } else {
        fallbackCopy(text, successMsg);
    }
}

function fallbackCopy(text, successMsg) {
    const textArea = document.createElement("textarea");
    textArea.value = text;
    textArea.style.position = "fixed";
    textArea.style.top = "-9999px";
    textArea.style.left = "-9999px";
    document.body.appendChild(textArea);
    textArea.focus();
    textArea.select();
    try {
        if (document.execCommand('copy')) alert(successMsg + ' (備用模式)');
        else alert('❌ 瀏覽器阻擋了自動複製。');
    } catch (err) {
        alert('❌ 發生錯誤，請手動複製。');
    }
    document.body.removeChild(textArea);
}

// 頁面載入時先執行一次渲染，確保狀態正確
document.addEventListener('DOMContentLoaded', () => {
    renderStagingArea();
});