function toggleButtonAbility(able) {
  if (able) {
    document.getElementById('btn-register').classList.remove('disabled');
  } else {
    document.getElementById('btn-register').classList.add('disabled');
  }
}

document.getElementById('btn-register').addEventListener('click', () => {
  toggleButtonAbility(false);

  const data = {
    username: document.getElementById('input-name').value,
    email: document.getElementById('input-email').value,
    password: document.getElementById('input-password').value,
  };
  if (data.password.length < 8) {
    alert('Password should be 8 letters or longer.');
    toggleButtonAbility(true);
    return;
  }

  post('/apis/v1/register', data)
    .then(() => {
      alert('A confirmation email has been sent. Please check your mailbox.');
      window.location.href = '/web/signin';
    })
    .catch(() => toggleButtonAbility(true));
});
