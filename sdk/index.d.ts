export interface LabangramClientOptions {
  baseUrl?: string;
  fetchImpl?: typeof fetch;
}

export interface InquiryDraft {
  status: 'requires_human_confirmation';
  message: string;
  contact_url: string;
  fields: {
    name: string;
    email: string;
    service_type: string;
    message: string;
  };
}

export class LabangramApiError extends Error {
  status: number;
  body: unknown;
}

export class LabangramClient {
  constructor(options?: LabangramClientOptions);
  request(path: string, init?: RequestInit): Promise<unknown>;
  getProjects(params?: Record<string, string | number>): Promise<unknown>;
  getProject(slug: string): Promise<unknown>;
  getServices(params?: Record<string, string>): Promise<unknown>;
  getPricing(currency?: 'TWD'): Promise<unknown>;
  prepareInquiry(input?: Partial<InquiryDraft['fields']>): InquiryDraft;
  callProductTool(name: string, args?: Record<string, unknown>): Promise<unknown>;
}

export const defaultClient: LabangramClient;
