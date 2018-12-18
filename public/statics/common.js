function post(url, body) {
  return new Promise((resolve, reject) => {
    axios.post(url, body)
      .then((res) => resolve(res.data))
      .catch((err) => {
        if (err.response) {
          alert(err.response.data.message);
        } else if (err.request) {
          console.error(err.request);
          alert('An error occurred. Please retry again.');
        } else {
          console.error(err.message);
          alert('An error occurred. Please retry again.');
        }
        reject();
      });
  });
}
