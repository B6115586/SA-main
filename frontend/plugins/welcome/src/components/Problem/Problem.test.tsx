import React from 'react';
import { render } from '@testing-library/react';
import ExampleComponent from './ExampleComponent';
import { ThemeProvider } from '@material-ui/core';
import { lightTheme } from '@backstage/theme';
import { rest } from 'msw';
import { setupServer } from 'msw/node';


describe('ExampleComponent', () => {
  const server = setupServer();
  // Enable API mocking before tests.
  beforeAll(() => server.listen({ onUnhandledRequest: 'error' }))

  // Reset any runtime request handlers we may add during the tests.
  afterEach(() => server.resetHandlers())

  // Disable API mocking after the tests are done.
  afterAll(() => server.close())

  // setup mock response
  beforeEach(() => {
    server.use(rest.get('/*', (_, res, ctx) => res(ctx.status(200), ctx.json({}))))
  })

  it('should render', () => {
    const rendered = render(
      <ThemeProvider theme={lightTheme}>
        <ExampleComponent />
      </ThemeProvider>,
      );
      expect(rendered.getByText('Welcome to welcome!')).toBeInTheDocument();
  });
});
