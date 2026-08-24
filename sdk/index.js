const DEFAULT_BASE_URL = 'https://labangram.kamera-ichi.com';
const PRODUCT_MCP_PATH = '/api/product-mcp';

export class LabangramApiError extends Error {
  constructor(message, status, body) {
    super(message);
    this.name = 'LabangramApiError';
    this.status = status;
    this.body = body;
  }
}

export class LabangramClient {
  constructor({ baseUrl = DEFAULT_BASE_URL, fetchImpl = globalThis.fetch } = {}) {
    if (typeof fetchImpl !== 'function') throw new TypeError('A fetch implementation is required.');
    this.baseUrl = baseUrl.replace(/\/$/, '');
    this.fetch = fetchImpl;
  }

  async request(path, init = {}) {
    const response = await this.fetch(new URL(path, this.baseUrl), {
      ...init,
      headers: {
        Accept: 'application/json',
        ...(init.headers || {}),
      },
    });
    const contentType = response.headers.get('content-type') || '';
    const body = contentType.includes('json') ? await response.json() : await response.text();
    if (!response.ok) {
      const detail = body && typeof body === 'object' && 'detail' in body ? body.detail : response.statusText;
      throw new LabangramApiError(String(detail || 'Labangram request failed.'), response.status, body);
    }
    return body;
  }

  getProjects(params = {}) {
    const query = new URLSearchParams(params);
    return this.request(`/api/v1/projects${query.toString() ? `?${query}` : ''}`);
  }

  getProject(slug) {
    if (!slug) throw new TypeError('A project slug is required.');
    return this.request(`/api/v1/projects/${encodeURIComponent(slug)}`);
  }

  getServices(params = {}) {
    const query = new URLSearchParams(params);
    return this.request(`/api/v1/services${query.toString() ? `?${query}` : ''}`);
  }

  getPricing(currency = 'TWD') {
    return this.request(`/api/v1/pricing?currency=${encodeURIComponent(currency)}`);
  }

  prepareInquiry({ name = '', email = '', message = '', service_type = '' } = {}) {
    return {
      status: 'requires_human_confirmation',
      message: 'Inquiry prepared; human confirmation is required before any submission.',
      contact_url: `${this.baseUrl}/about#contact`,
      fields: { name, email, service_type, message },
    };
  }

  async callProductTool(name, args = {}) {
    return this.request(PRODUCT_MCP_PATH, {
      method: 'POST',
      headers: {
        Accept: 'application/json, text/event-stream',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        jsonrpc: '2.0',
        id: Date.now().toString(36),
        method: 'tools/call',
        params: { name, arguments: args },
      }),
    });
  }
}

export const defaultClient = new LabangramClient();
