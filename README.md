Team admin ingests all the documents from knowledge base, including application’s historical data
Initial bot was tried with pdf documents, then extended to txt, doc
Current bot accepts below formats for file ingestions
Doc, docx, pdf, hanword
Ppt, csv,
Markdown files, mails, html files
Json, image files
Video, audio
Using various splitters available in langchain/llamaindex, with the chunk size provided, sequence of transformed documents will be returned
Based on the pre-trained embedding model used, transformerembeddings will compute embeddings
These embeddings will then be stored on to vector database





This AI bot uses Retrieval-Augmented Generation (RAG) approach which combines response retrieval specific to internal knowledge base, internal data and with generative models
User submits the query from UI, which would be converted to embedding using the pre-trained embedding model and this embedding is then used to perform similarity search.
Before retrieval from vectordb, LLM would not be used, but used only after retrieval, to synthesize the response using original query and contextual information 
Retriever (VectorIndexRetriever) from llamaindex/langchain is used to search from embedding stored in vector database
Similarity top-k search would be done (ranked by similarity by retrieving top-k results)
Retrieved top-k results will then be passed to LLM context which generates more accurate and relevant response
Based on the prompt provided by user, response will be rephrased, summarized accordingly by LLM



SOPs ingested to bot, if someone new to team wants to know the process ex: datafix executions process, can query and get required answer wrt to process being followed
Pushing the server status if the services are up or down on specific time interval, there by getting overall up or down time and predict any down times
Ingesting api response times and get required details
Uploading all the server details of preprod regions and query for what's needed




Data driven decision making
Can analyse past interactions, user feed back based on the prompts
Predictive analytics
Predict potential issues, down time for proactive measures
Personalizations
AI can provide insights into user behavior and preferences for personalized interactions
Feedback loop integration
AI can automate the feedback loop by continuously learning from user interactions and adjusting the retrieval and generation parameters accordingly.

Future use cases: 
Application downtime and spike or drop in the number of calls predictions based on the application logs, CLS, historical data from AppD, DataDog, ServiceNow, root causes repository.
Anomaly detection on platform.
Analyze incidents, issue wrt network outage, performance and predict future occurrences.
Automated remediation actions such as flipping, restart, configuration changes based on thresh hold values set.
Forecast the server metrics based on the traffic.
For risks, vulnerabilities detection and remediating



Running on CPU is very slow and is having huge latency issues
Inefficient caching large parts of model on CPU memory
Handling large datasets
No huge model's availability within nexus and downloading them also takes more time.
With a midsized model not very detailed embedding available
Current model is ~ 4GB in size
Tried with various vector dbs but dependency issues with other packages when running on a CPU
Dependency on various packages not available on nexus, not compatible with python versions which we currently have



import React, { useState } from 'react';

const FileDeployList = () => {
  const [filename, setFilename] = useState('');
  const [patchList, setPatchList] = useState([]);
  const [error, setError] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch(`http://localhost:8080/patch-list?filename=${filename}`);
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data = await response.json();
      setPatchList(data);
      setError(null);
    } catch (error) {
      setError('Failed to fetch patch list.');
      console.error(error);
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label>
          Filename:
          <input type="text" value={filename} onChange={(e) => setFilename(e.target.value)} />
        </label>
        <button type="submit">Get File Deploy List</button>
      </form>

      {error && <p style={{ color: 'red' }}>{error}</p>}

      {patchList.length > 0 && (
        <div>
          <h2>File Deploy List:</h2>
          <table border="1">
            <thead>
              <tr>
                <th>Deploy File</th>
                <th>Source File Size</th>
                <th>Patch ID</th>
                <th>Release ID</th>
                <th>Modification Time</th>
              </tr>
            </thead>
            <tbody>
              {patchList.map((patch) => (
                <tr key={patch.patch_id}>
                  <td>{patch.deploy_file}</td>
                  <td>{patch.src_file_size}</td>
                  <td>{patch.patch_id}</td>
                  <td>{patch.release_id}</td>
                  <td>{patch.mod_time}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
};

export default FileDeployList;
