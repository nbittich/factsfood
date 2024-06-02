
document.addEventListener("DOMContentLoaded", () => {
  const modal = document.getElementById('scanModal');
  const qrCodeInput = document.getElementById("qrCodeInput");
  const qrScanBtn = document.getElementById("qrScanModalBtn");
  const scanner = new Html5Qrcode("reader");
  const config = { fps: 10, };

  modal.addEventListener('shown.bs.modal', () => {
    scanner.start({ facingMode: "environment" }, config, onScanSuccess);
  })

  modal.addEventListener('hide.bs.modal', async () => {
    await scanner.stop();
  })

  function onScanSuccess(decodedText, decodedResult) {
    console.log(`Code scanned = ${decodedText}`, decodedResult);
    qrCodeInput.value = decodedText;
    toggleModal('scanModal', false);
  }


  qrScanBtn.onclick = (e) => {
    e.stopPropagation();
    e.preventDefault();
    toggleModal('scanModal', true);
  };
})
