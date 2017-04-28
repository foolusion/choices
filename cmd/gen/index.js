javascript:(function() {
  fetch("http://gen.ato.platform.prod.aws.cloud.nordstrom.net/").then(r =>
    r.json().then(data => {
      let experiments = data.experiments.map(experiment => {
        let params = experiment.params.map(param => {
          let choices =  param.value.choices.map(choice => {
            let button = document.createElement("button");
            button.textContent = `${choice}`;
            button.addEventListener("click", ev => {
              fetch(`http://gen.ato.platform.prod.aws.cloud.nordstrom.net/experiment/${experiment.id}/param/${param.name}/value/${choice}`).then(response => 
                response.json()
              ).then(data => {
                document.cookie = `experiments=ExperimentId=${data};domain=.nordstrom.com;path=/`;
                location.reload();
              });
            });
            return button;
          });
          let div = document.createElement("div");
          let paramName = document.createElement("span");
          paramName.textContent = `Param: ${param.name}`;
          div.appendChild(paramName);
          choices.forEach(b => div.appendChild(b));
          return div;
        })
        let div = document.createElement("div");
        let experimentName = document.createElement("span");
        experimentName.textContent = `Experiment: ${experiment.name}`;
        div.appendChild(experimentName);
        params.forEach(b => div.appendChild(b));
        div.style.margin = "16px 16px 0 16px";
        let labels = document.createElement("div");
        Object.keys(experiment.labels).forEach(k => {
          let l = document.createElement("span");
          l.textContent = `${k}: ${experiment.labels[k]}`;
          labels.appendChild(l);
          labels.appendChild(document.createElement("br"));
        });
        div.appendChild(labels);
        return div;
      })
      let div = document.createElement("div");
      experiments.forEach(b => {
        div.appendChild(b);
      })
      div.style.position = "absolute";
      div.style.zIndex = "10";
      div.style.background = "orange";
      div.style.fontSize = "20px";
      document.body.insertBefore(div, document.body.firstChild);
    })
  )
})();