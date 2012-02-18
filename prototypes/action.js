var Action = function (parameters, precondition, effect) {
  this.parameters = parameters;
  this.precondition = precondition;
  this.effect = effect;
}

Action.prototype = Object.create(Object);
Action.prototype.constructor = Action;

Action.prototype.act = function (state) { 
}

Action.prototype.validFor = function (state) {
}
