logic = require('./predicate.js')

var CPProblem = function(predicates, actions, initial, end) {
  this.state = {};
  this.predicates = predicates;
  this.actions = actions;
  this.initial = initial;
  this.end = end;
}

CPProblem.prototype = Object.create(Object);
CPProblem.prototype.constructor = CPProblem;

function buildPredicates (predicatesData) {
  
  var predicates = {}
  
  for (var key in predicatesData) {
    var p = new logic.Predicate(key, predicatesData[key])
    predicates[key] = p;
  }

  return predicates;
}

function buildActions (actionsData) {
  for (var key in actionsData) {
    var a = new logic.Action()
  }
}

var buildProblem = function (data) {
  var predicates = buildPredicates(data.predicates);

  return new CPProblem(predicates)


}

module.exports.buildProblem = buildProblem;
module.exports.CPProblem = CPProblem;
