
function toggleModal(id, show) {
  const modalElt = document.getElementById(id);
  let modal;
  if (show) {
    modal = new bootstrap.Modal(modalElt);
    modal.show();
  } else {
    modal = bootstrap.Modal.getInstance(modalElt)
    modal.hide();
  }
  return modal;
}

function onSubmitForm() {
  for (const btn of document.querySelectorAll("[type='submit']")) {
    btn.disabled = true;
  }
  return true
}
