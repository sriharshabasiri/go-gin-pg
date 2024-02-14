 {/* Form for creating a new refresh entry */}
      <form onSubmit={handleSubmit}>
        <label>
          Environment:
          <select
            name="environment"
            value={newRefresh.environment}
            onChange={handleInputChange}
            required
          >
            {/* Options for Environment dropdown */}
            <option value="production">Production</option>
            <option value="testing">Testing</option>
            <option value="development">Development</option>
            {/* Add more options as needed */}
          </select>
        </label>

        <label>
          Code Base:
          <input
            type="text"
            name="code_base"
            value={newRefresh.code_base}
            onChange={handleInputChange}
            required
          />
        </label>

        <label>
          Change Ticket:
          <input
            type="text"
            name="change_ticket"
            value={newRefresh.change_ticket}
            onChange={handleInputChange}
            required
          />
        </label>

        <button type="submit">Create</button>
      </form>
