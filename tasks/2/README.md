
# Candidate Task 2

Using the Sensor Simulator API available in the [GitHub repository](https://github.com/barisvelioglu/candidate-level-1), your task is to:

1. **Retrieve sensor data:** Collect data from the `/temperature`, `/humidity`, and `/vibration` endpoints every second.
2. **Enrich the data:** Add a new field, `retrieved_timestamp`, which should store the exact timestamp of when the data was fetched from the API.
3. **Store the data:** Save the enriched data into a persistent storage solution of your choice. Ensure there is no data loss.
4. **Apply a retention policy:** Ensure the stored data is automatically deleted after exactly **1 day**.

---

## Expectations

1. **Reliability:** The solution should guarantee that no sensor data is lost during the process.
2. **Efficiency:** Implement your solution in a way that avoids excessive resource usage or unnecessary API requests.
3. **Retention management:** Ensure data older than 24 hours is cleaned up automatically using a process or mechanism that fits your chosen storage solution.

---

### **Guidelines**
- Your solution should be able to handle unexpected failures (e.g., API downtime or storage unavailability) gracefully.
- The API and simulator details can be found in the GitHub repository.
- You are free to use any tools, technologies, or libraries you prefer, but be prepared to explain your choices.

Good luck!