import { api, Tag } from '@/api';
import {
  AddStudent,
  GetStudentDetailProps,
  GetTeachers,
  ReviewStudentStatusRequest,
  StudentFilter,
  StudentProps,
  StudentPropsWithId
} from '../types';
import { getQueryString } from '@/utils/helpers/get-query-string';
import { UserAccountBasicProps } from '@/components/user-account-basic';

export const studentApi = api.injectEndpoints({
  endpoints: (builder) => ({
    getStudents: builder.query<UserAccountBasicProps[], StudentFilter>({
      query: (payload) => {
        const queryString = getQueryString(payload);
        return `/students${queryString}`;
      },
      providesTags: (result) =>
        result?.map(({ id , name,email, role,systemAccess, lastLogin}) => {
          return { type: Tag.STUDENTS, id, name, email, role, systemAccess, lastLogin };
        }) || [{ type: Tag.STUDENTS }]
    }),
    getStudentDetail: builder.query<GetStudentDetailProps, string | undefined>({
      query: (id) => (id ? `/students/${id}` : `/account/me`),
      providesTags: (result) => (result ? [{ type: Tag.STUDENTS, id: result.id }] : [])
    }),
    reviewStudentStatus: builder.mutation<{ message: string }, ReviewStudentStatusRequest>({
      query: ({ id, status }) => ({
        url: `/students/${id}/status`,
        method: 'POST',
        body: { status }
      }),
      invalidatesTags: (_result, _error, { id }) => [{ type: Tag.STUDENTS, id }]
    }),
    addStudent: builder.mutation<AddStudent, StudentProps>({
      query: (payload) => ({
        url: `/students`,
        method: 'POST',
        body: payload
      }),
      invalidatesTags: [Tag.STUDENTS]
    }),
    updateStudent: builder.mutation<{ message: string }, StudentPropsWithId>({
      query: ({ id, ...payload }) => ({
        url: `/students/${id}`,
        method: 'PUT',
        body: payload
      }),
      invalidatesTags: (_result, _error, { id }) => [{ type: Tag.STUDENTS, id }]
    }),
    getTeachers: builder.query<GetTeachers, void>({
      query: () => `/teachers`
    })
  })
});

export const {
  useGetStudentsQuery,
  useLazyGetStudentDetailQuery,
  useReviewStudentStatusMutation,
  useAddStudentMutation,
  useUpdateStudentMutation,
  useGetTeachersQuery
} = studentApi;
