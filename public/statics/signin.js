function onGoogleSuccess(googleUser) {
  const params = new URLSearchParams(window.location.search);
  const data = {
    appId: params.get('appId'),
    provider: 'google',
    payload: googleUser.getAuthResponse().id_token,
  };

  axios.post('/apis/signin', data).then((res) => {
    console.log(res.data);
  }).catch((e) => {
    console.error(e.response.data);
  });
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
