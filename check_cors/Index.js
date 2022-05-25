   fetch("http://0.0.0.0:5000/api/v1/directions", {
      headers: { Authorization: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDYxNDU5MDUsImlhdCI6MTY0NjEzMDkwNSwibG9naW4iOiJjb29yZHRlc3RAYW5kZXJzZW5sYWIuY29tIiwicm9sZSI6ImNvb3JkaW5hdG9yIiwiaWQiOjE0fQ.XLDEBv5f0iiaQPad4z3ei2kZoxbS7zrsJU1rS4LGx0w", accept: "application/json" },
   })
      .then((r) => r.json())
      .then((r) => console.log(r));
