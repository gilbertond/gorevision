'use strict';

//Handler for path root. i.e. /
module.exports.home = async (event) => {
  return {
    statusCode: 200,
    body: JSON.stringify({
      message: 'You wanted the home page',
      // input: event,
    }, null, 2),
  };

  // Use this code if you don't use the http event with the LAMBDA-PROXY integration
  // return { message: 'Go Serverless v1.0! Your function executed successfully!', event };
};

module.exports.aboutGilbert = async (event) => {
  return {
    statusCode: 200,
    body: JSON.stringify({
      message: 'You asked about Gilbert!',
      // input: event,
    }, null, 2),
  };

  // Use this code if you don't use the http event with the LAMBDA-PROXY integration
  // return { message: 'Go Serverless v1.0! Your function executed successfully!', event };
};