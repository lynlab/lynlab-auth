function signin(data) {
  return axios.post('/apis/v1/signin', data).then((res) => {
    const params = new URLSearchParams(window.location.search);
    const url = `${params.get('redirectUrl')}?accessToken=${res.data.accessToken}&refreshToken=${res.data.refreshToken}&expireAt=${res.data.expireAt}`;
    window.location.replace(url);
  }).catch((e) => {
    switch (e.response.data.message) {
      case 'no_such_account':
        showRegisterModal(data);
        break;
      case 'authorization_required':
        showAuthorizationModal(data);
        break;
    }
  });
}

function register(data) {
  return axios.post('/apis/v1/register', data).then(() => {
    hideRegisterModal();
    signin(data);
  });
}

function authorize(data) {
  return axios.post('/apis/v1/authorize', data).then(() => {
    hideAuthorizationModal();
    signin(data);
  });
}

function showRegisterModal(data) {
  const modal = document.getElementById('register-modal');
  modal.classList.add('visible');

  document.getElementById('register-button').addEventListener('click', () => {
    const input = document.getElementById('input-username');
    const username = input.value;
    if (username.length === 0 || username.length > 20) {
      input.classList.add('error');
    } else {
      data.username = username;
      register(data);
    }
  });

  document.getElementById('close-register-modal').addEventListener('click', () => {
    hideRegisterModal();
  });
}

function hideRegisterModal() {
  const modal = document.getElementById('register-modal');
  modal.classList.remove('visible');
}

function showAuthorizationModal(data) {
  const modal = document.getElementById('authorization-modal');
  modal.classList.add('visible');

  document.getElementById('authorize-button').addEventListener('click', () => {
    authorize(data);
  });

  document.getElementById('close-authorization-modal').addEventListener('click', () => {
    modal.classList.remove('visible');
  });
}

function hideAuthorizationModal(data) {
  const modal = document.getElementById('register-modal');
  modal.classList.remove('visible');
}

/// Google Signin
function onGoogleSuccess(googleUser) {
  const params = new URLSearchParams(window.location.search);
  const data = {
    appId: params.get('appId'),
    provider: 'google',
    payload: googleUser.getAuthResponse().id_token,
  };

  signin(data);
}

function onGoogleFailure(err) {
  console.log(err);
}

function renderGoogleButton() {
  gapi.signin2.render('signin-google', {
    'scope': 'profile email',
    'width': 360,
    'height': 50,
    'longtitle': true,
    'theme': 'dark',
    'onsuccess': onGoogleSuccess,
    'onfailure': onGoogleFailure
  });
}
