function toggleButtonAbility(able) {
  if (able) {
    document.getElementById('btn-signin').classList.remove('disabled');
  } else {
    document.getElementById('btn-signin').classList.add('disabled');
  }
}

document.getElementById('btn-signin').addEventListener('click', () => {
  toggleButtonAbility(false);

  const data = {
    email: document.getElementById('input-email').value,
    password: document.getElementById('input-password').value,
  };

  post('/apis/v1/token/generate', data)
    .then((res) => {
      const params = new URLSearchParams(window.location.search);
      window.location.href = `${params.get('redirect_url')}?access_token=${res.access_token}&refresh_token=${res.refresh_token}`;
    })
    .catch(() => toggleButtonAbility(true));
});
