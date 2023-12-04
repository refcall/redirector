<!DOCTYPE html>
<html>
<head>
  <title>{{ .Config.Name }} redirection</title>
  <script>
    window.addEventListener('DOMContentLoaded', (event) => {
      setTimeout(() => {
        document.querySelector('.progress-content').style.width = '100%'
      }, 200)
      setTimeout(() => {
        window.location = "/redirect";
      }, {{ .Config.RedirectSeconds }} * 1000)
    })
  </script>
  <style>
    html, body {
      margin: 0;
      padding: 0;
      background-color: {{ .Config.BackgroundColor }};
      height: 100%;
      font-family: Arial, Helvetica, sans-serif;
    }
    .container {
      height: 100%;
      display: flex;
      align-items: center;
      justify-content: center;
    }
    .card {
      background-color: {{ .Config.CardColor }};
      max-width: 400px;
      padding: 40px;
      border-radius: 60px;
      box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
      display: flex;
      flex-direction: column;
    }
    .card .image {
      height: 50px;
      margin-bottom: 1.8rem;
      object-fit: contain;
    }
    .card .text {
      margin-bottom: 2rem;
    }
    .progress {
      position: relative;
      width: 100%;
      height: 1rem;
      background-color: {{ .Config.ProgressBackgroundColor }};
      border-radius: 60px;
      margin-bottom: 2rem;
    }
    .progress-content {
      position: absolute;
      height: 100%;
      width: 4%;
      background-color: {{ .Config.ProgressColor }};
      border-radius: 60px;
      transition: width {{ .Config.ProgressSeconds }}s;
    }
  </style>
</head>
<body>
  <div class="container">
    <div class="card">
      <img class="image" src="{{ .Config.Logo }}" />
      <span class="text">{{ .Translation.Text }}</span>
      <div class="progress">
        <div class="progress-content"></div>
      </div>
      <a class="redirect" href="/redirect">{{ .Translation.Redirect }}</a>
    </div>
  </div>
</body>
</html> 