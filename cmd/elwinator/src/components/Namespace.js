import React from 'react';
import { Link } from 'react-router';
import { connect } from 'react-redux';

import NavSection from './NavSection';
import LabelList from './LabelList';
import NewLabel from './NewLabel';
import Segment from './Segment';
import ExperimentList from './ExperimentList';
import { labelNewURL, experimentNewURL } from '../urls';
import { addLabel, toggleLabel } from '../actions';

const Namespace = ({ ns, addLabel, toggleLabel }) => {
  const llProps = { namespaceName: ns.name, labels: ns.labels, toggleLabel };
  const newLabelProps = { namespaceName: ns.name, addLabel, redirectOnSubmit: false };
  const nsSegments = ns.experiments.reduce((prev, e) => {
      e.segments.forEach((seg, i) => {
        prev[i] |= seg;
      });
      return prev;
    }, new Uint8Array(16).fill(0));
  return (
    <div className="container">
      <div className="row"><h1>{ ns.name }</h1></div>
      <div className="row">
        <NavSection>
          <Link to={ labelNewURL(ns.name) } className="nav-link">New label</Link>
          <Link to={ experimentNewURL(ns.name) } className="nav-link">New Experiment</Link>
        </NavSection>
        <div className="col-sm-9">
          <h2>Labels</h2>
          <LabelList { ...llProps } />
          <NewLabel { ...newLabelProps } />
          <h2>Segments</h2>
          <Segment namespaceSegments={nsSegments} />
          <h2>Experiments</h2>
          <ExperimentList namespaceName={ ns.name } />
        </div>
      </div>
    </div>
  );
};

const mapStateToProps = (state, ownProps) => {
  const ns = state.namespaces.find(n => n.name === ownProps.params.namespace);
  return {
    ns,
  }
};

const mapDispatchToProps = {
  addLabel,
  toggleLabel,
}

const connected = connect(mapStateToProps, mapDispatchToProps)(Namespace);
 
export default connected;