
function openModal(id) {
  document.addEventListener("DOMContentLoaded", function () {
    var modal = new bootstrap.Modal(document.getElementById(id));
    modal.show();
  });
}

function onSubmitForm() {
  for (const btn of document.querySelectorAll("[type='submit']")) {
    btn.disabled = true;
  }
  return true
}
